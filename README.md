# ASP Contracts V1.0

> [!CAUTION]
> These Smart Contracts are pending audit & should be used with caution.

This repository contains the Solidity smart contracts utilised by the "association-set provider" ASP with go & ts packages for interacting with the contracts.

Please refer to the [ASP Spec V1.0](https://0xbow-io.github.io/asp-spec-v1.0/) for more context on the "association set provider" (ASP).

---

## Current Version

1.0.0-alpha

## Version History

| Version     | Release Date | Changes            |
| ----------- | ------------ | ------------------ |
| 1.0.0-alpha | 2024-07-22   | Core ASP Contracts |

## Planned Future Changes

- N/A

## Deprecation Notices

- N/A

## Current Deployments

| Chain   | Contract Address                                |
| ------- | ----------------------------------------------- |
| Sepolia | [0xaaf9221e9c337ba5ba2c4bd6c0af8b2d2b9cd286][1] |
| Gnosis  | [0xaaf9221e9c337ba5ba2c4bd6c0af8b2d2b9cd286][2] |
| Mainnet | ⚠️ Pending Deployment                           |
| Base    | ⚠️ Pending Deployment                           |

[1]: https://sepolia.etherscan.io/address/0xaaf9221e9c337ba5ba2c4bd6c0af8b2d2b9cd286
[2]: https://gnosisscan.io/address/0x9a6425a02a72C6983E45A4939E432Ebe4611f53E

## 1. Development:

### 1.1. Setup:

Ensure you have the following installed:

- [ ] [Foundry Forge (>=v0.2.0)](https://book.getfoundry.sh/getting-started/installation)
- [ ] [Golang (==v1.22.5)](https://go.dev/doc/install)
- [ ] [AbiGen (>=1.1.4.7-stable](https://geth.ethereum.org/docs/tools/abigen)
- [ ] [Bun (>=v1.1.19)](https://bun.sh/)


### 1.2. Testing:

- To run the smart-contracts unit-tests:

```bash
forge test -vvvv
```

This will run the tests in the `test` directory:

- `AssociationSetProvider.test.ts`
- `AssociationSet.test.ts`

### 1.3. Contract Deployment:

```bash
## Set these environment variables in the `.env` file:

# RPC_URL="..."
# PRIVATE_KEY="..."
# ETHERSCAN_API_KEY="..."
# Chain="..."

## Load the environment variables:
source .env

## To deploy AssociationSetProvider.sol:

forge script --chain $CHAIN script/deployment.s.sol:Deploy_AssociationSetProvider --rpc-url $RPC_URL --legacy --broadcast --verify -vvv
```

---

## 2. Integration:

Refer to [`integrations`](/integrations/README.md) for integrating the ASP with other systems / protocols.

---

## 3. How to Contribute:

---

## 4. Versioning:

### 4.1 Backwards Compatibility

ASP Contracts V1.0 establishes the baseline for future versions. All subsequent minor versions (1.x) will maintain backwards compatibility with V1.0.

### 4.2 Semantic Versioning

ASP Contracts follows Semantic Versioning 2.0.0 (https://semver.org/). Version numbers are in the format MAJOR.MINOR.PATCH, where:

- MAJOR version increments denote incompatible API changes
- MINOR version increments add functionality in a backwards-compatible manner
- PATCH version increments make backwards-compatible bug fixes
- PRE-RELEASE versions are denoted by appending a hyphen and a series of dot-separated identifiers immediately following the PATCH version
