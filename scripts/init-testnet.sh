#!/bin/sh
source .env

VALIDATOR_NAME=validator1
CHAIN_ID=opkit
KEY_NAME=opkit-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

DA_BLOCK_HEIGHT=$(curl https://rpc-mocha.pops.one/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT
echo $DA_ADDRESS

# build the opkit chain with Rollkit
ignite chain build

# reset any existing genesis/chain data
opkitd tendermint unsafe-reset-all

# initialize the validator with the chain ID you set
opkitd init $VALIDATOR_NAME --chain-id $CHAIN_ID

# add keys for key 1 to keyring-backend test
opkitd keys add $KEY_NAME --keyring-backend test

# add these as genesis accounts
opkitd genesis add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test

# set the staking amounts in the genesis transaction
opkitd genesis gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test

# collect genesis transactions
opkitd genesis collect-gentxs

# copy centralized sequencer address into genesis.json
# Note: validator and sequencer are used interchangeably here
ADDRESS=$(jq -r '.address' ~/.opkit/config/priv_validator_key.json)
PUB_KEY=$(jq -r '.pub_key' ~/.opkit/config/priv_validator_key.json)
jq --argjson pubKey "$PUB_KEY" '.consensus["validators"]=[{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1000", "name": "Rollkit Sequencer"}]' ~/.opkit/config/genesis.json > temp.json && mv temp.json ~/.opkit/config/genesis.json

opkitd start \
    --rollkit.aggregator \
    --rollkit.da_address $DA_ADDRESS \
    --rollkit.da_auth_token $DA_AUTH_TOKEN \
    --rollkit.da_namespace $DA_NAMESPACE \
    --rollkit.da_start_height $DA_BLOCK_HEIGHT \
    --minimum-gas-prices="0.025stake"
