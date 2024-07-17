##!/bin/bash

## Get all the dependencies
echo "Fetching dependencies..."
forge install OpenZeppelin/openzeppelin-contracts
forge install privacy-scaling-explorations/zk-kit.solidity
forge install vimwitch/poseidon-solidity


## generate the Ts & Go ABI bindings
## echo "generating ABI bindings..."
## bun run ./gen_ts_abi_bindings.ts
