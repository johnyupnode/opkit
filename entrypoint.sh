#!/bin/sh

VALIDATOR_NAME=validator1
KEY_NAME=${CHAIN_ID}-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

DA_BLOCK_HEIGHT=$(curl ${RPC_URL}/block | jq -r '.result.block.header.height')
echo "Block Height:" $DA_BLOCK_HEIGHT
echo "DA Address:" $DA_ADDRESS

# build the chain with Rollkit
ignite chain build

# Symbolic link to daemon
ln -s $(which ${CHAIN_ID}d) /usr/local/bin/daemon

# reset any existing genesis/chain data
daemon tendermint unsafe-reset-all

# initialize the validator with the chain ID you set
daemon init $VALIDATOR_NAME --chain-id $CHAIN_ID

# add keys for key 1 to keyring-backend test
daemon keys add $KEY_NAME --keyring-backend test

# add these as genesis accounts
daemon genesis add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test

# set the staking amounts in the genesis transaction
daemon genesis gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test

# collect genesis transactions
daemon genesis collect-gentxs

# copy centralized sequencer address into genesis.json
# Note: validator and sequencer are used interchangeably here
ADDRESS=$(jq -r '.address' ~/.${CHAIN_ID}/config/priv_validator_key.json)
PUB_KEY=$(jq -r '.pub_key' ~/.${CHAIN_ID}/config/priv_validator_key.json)
jq --argjson pubKey "$PUB_KEY" '.consensus["validators"]=[{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1000", "name": "Rollkit Sequencer"}]' ~/.${CHAIN_ID}/config/genesis.json > temp.json && mv temp.json ~/.${CHAIN_ID}/config/genesis.json

# Start the chain
daemon start \
  --rollkit.aggregator \
  --rollkit.da_address $DA_ADDRESS \
  --rollkit.da_auth_token $DA_AUTH_TOKEN \
  --rollkit.da_namespace $DA_NAMESPACE \
  --rollkit.da_start_height $DA_BLOCK_HEIGHT \
  --minimum-gas-prices="0.025stake"