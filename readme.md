# opkit
**opkit** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/opkit@latest! | sudo bash
```
`username/opkit` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## How to use

### Prerequisite

The domain module need to be funded with stake token to be able to distribute the domain minting reward

Domain module: opkit17tlc8ps2fhpq8xyw6x3zhg0jzgmlqj4a7xeflh

### CLI

The blockchain node binary can be used to interact with the blockchain. To see the available commands, run:

#### List all domain by opkit address
```shell
opkitd query domain list-domain-opkit <opkit_address>
```
Example:
```shell
opkitd query domain list-domain-opkit opkit1t85skec9sf386946y8g820auhpxw8z7csky8j5

Domain:
- domain: johny.opkit
  owner: 0xc4C565bdF45F54cca35036347F8BC1974c800372
  string_records:
  - key: opkit
    value: opkit1t85skec9sf386946y8g820auhpxw8z7csky8j5
```

#### List all domain by evm address
```shell
opkitd query domain list-domain-evm <owner_address>
```
Example:
```shell
opkitd query domain list-domain-evm 0xc4C565bdF45F54cca35036347F8BC1974c800372

Domain:
- domain: test2.opkit
  owner: 0xc4C565bdF45F54cca35036347F8BC1974c800372
  string_records:
  - key: github
    value: EZ420
  - key: opkit
    value: opkit13ng4lj30e6ptpv8gm5rf2dqqyzxkm5tskvtmal
- domain: johny.opkit
  owner: 0xc4C565bdF45F54cca35036347F8BC1974c800372
  string_records:
  - key: opkit
    value: opkit1t85skec9sf386946y8g820auhpxw8z7csky8j5
```

#### List all domain by key and value of string record
```shell
opkitd query domain list-domain-by-string <key> <value>
```
Example:
```shell
opkitd query domain list-domain-by-string github EZ420

Domain:
- domain: test2.opkit
  owner: 0xc4C565bdF45F54cca35036347F8BC1974c800372
  string_records:
  - key: github
    value: EZ420
  - key: opkit
    value: opkit13ng4lj30e6ptpv8gm5rf2dqqyzxkm5tskvtmal
```

#### Set primary domain
```shell
opkitd tx domain set-primary-domain <domain> --from opkit-key --fees 5000stake
```
Example:
```shell
opkitd tx domain set-primary-domain test2.opkit --from opkit-key --fees 5000stake
```

#### Get primary domain
```shell
opkitd query domain get-primary-domain <domain>
```
Example:
```shell
opkitd query domain get-primary-domain johny.opkit

Domain:
- domain: johny.opkit
  owner: 0xc4C565bdF45F54cca35036347F8BC1974c800372
  string_records:
  - key: opkit
    value: opkit1t85skec9sf386946y8g820auhpxw8z7csky8j5
```

#### Claim reward for domain
```shell
opkitd tx domain claim-reward <domain> --from opkit-key --fees 5000stake
```
Example:
```shell
opkitd tx domain claim-reward johny.opkit --from opkit-key --fees 5000stake
```

#### Query reward for domain
```shell
opkitd query domain reward <domain>
```
Example:
```shell
opkitd query domain reward johny.opkit 

Reward:
  amount:
    amount: "100000000"
    denom: stake
  domain: johny.opkit
  is_claimed: false
```


## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
