##!/bin/bash

## Get all the dependencies for smart contracts
echo "Fetching dependencies..."
rm -rf lib
forge install OpenZeppelin/openzeppelin-contracts privacy-scaling-explorations/zk-kit.solidity vimwitch/poseidon-solidity
forge remappings
