// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/AssociationSetProvider.sol";

contract Deploy_AssociationSetProvider is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        // Deploy the AssociationSetProvider contract
        new AssociationSetProvider();
        vm.stopBroadcast();
    }
}
