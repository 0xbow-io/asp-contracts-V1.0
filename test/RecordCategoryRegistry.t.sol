// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

import "../src/RecordCategoryRegistry.sol";

/**
 * @title TestRecordCategoryRegistry
 * @dev Test the RecordCategoryRegistry contract
 */
contract TestRecordCategoryRegistry is Test {
    RecordCategoryRegistry internal registry;

    function setUp() public {
        registry = new RecordCategoryRegistry(address(this));
    }

    /**
     * @notice The registry admin role should only be granted to the contract admin
     * @dev Test granting the registry admin role:
     * 1. Grant the registry admin role to address 0x01
     * 2. Attempt to grant the registry admin role to address 0x01 without the necessary permissions
     */
    function test_GrantRegistryAdminRole() public {
        bytes32 adminRole = registry.REGISTRY_ADMIN_ROLE();

        /// Grant the registry admin role to address 0x01
        registry.grantRole(adminRole, address(0x01));
        assertTrue(registry.hasRole(adminRole, address(0x01)));

        /// Attempt to grant the registry admin role to address 0x01 without the necessary permissions
        vm.startPrank(address(0x2));
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                0x0000000000000000000000000000000000000002,
                0x0000000000000000000000000000000000000000000000000000000000000000
            )
        );
        registry.grantRole(adminRole, address(0x01));
    }

    /**
     * @notice The registry admin role should only be revoked by the contract admin
     * @dev Test revoking the registry admin role:
     * 1. Grant the registry admin role to address 0x01
     * 2. Revoke the registry admin role from address 0x01
     */
    function test_RevokeRegistryAdminRole() public {
        bytes32 adminRole = registry.REGISTRY_ADMIN_ROLE();

        /// Grant the registry admin role to address 0x01
        registry.grantRole(adminRole, address(0x01));
        assertTrue(registry.hasRole(adminRole, address(0x01)));

        /// Revoke the registry admin role from address 0x01
        registry.revokeRole(adminRole, address(0x01));
        assertFalse(registry.hasRole(adminRole, address(0x01)));
    }

    /**
     * @notice The registry admin role should only be allowed to
     *         set the category for a record
     * @dev Test setting the category for a record:
     * 1. Set the category for a record
     * 2. Attempt to set the category for a record without the necessary permissions
     */
    function test_SetCategoryForRecord() public {
        /// Set the category for a record
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(1)),
            bytes32(uint256(2))
        );
        (
            uint256 root,
            bytes32 recordHash,
            bytes32 categoryBitmap,
            uint256 index
        ) = registry.getLatestForScope(0x01);

        /// Set the category for another record
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(3)),
            bytes32(uint256(4))
        );

        (
            uint256 NewRoot,
            bytes32 NewRecordHash,
            bytes32 NewCategoryBitmap,
            uint256 NewIndex
        ) = registry.getLatestForScope(0x01);

        /// expect diffs
        assertTrue(root != NewRoot);
        assertTrue(recordHash != NewRecordHash);
        assertTrue(categoryBitmap != NewCategoryBitmap);
        assertTrue(index != NewIndex);

        /// Attempt to set the category for a record without the necessary permissions
        vm.startPrank(address(0x2));
        vm.expectRevert(
            abi.encodeWithSelector(
                RecordCategoryRegistry.InsufficientRole.selector,
                address(0x02),
                bytes32(
                    0xbb28eb1a0cfabcecf96003fab466159bc2e051e49d79baf049890044e9072440
                )
            )
        );
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(4)),
            bytes32(uint256(5))
        );
    }

    /**
     * @notice Testing the getCategoryBitmap function
     * @dev Test getting the category bitmap for a record:=
     */
    function test_GetCategoryBitmap() public {
        /// Set the category for a record
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(1)),
            bytes32(uint256(2))
        );

        bytes32 categoryBitMap = registry.getCategoryBitmap(
            0x01,
            bytes32(uint256(1))
        );
        assertTrue(categoryBitMap == bytes32(uint256(2)));
    }

    /**
     * @notice Testing the getRecordHashAndCategoryAt function
     * @dev Test getting the record hash and category bitmap for a record at a specific index
     */
    function test_GetRecordHashAndCategoryAt() public {
        /// Set the category for a record
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(1)),
            bytes32(uint256(2))
        );

        (bytes32 recordHash, bytes32 categoryBitMap) = registry
            .getRecordHashAndCategoryAt(0x01, 0);
        assertTrue(recordHash == bytes32(uint256(1)));
        assertTrue(categoryBitMap == bytes32(uint256(2)));
    }

    /**
     * @notice Testing the getRecordHashAndCategoryAt function
     * @dev Test getting the record hash and category bitmap for a record at a specific index
     */
    function test_TryGetCategoryBitmap() public {
        /// Set the category for a record
        registry.setCategoryForRecord(
            0x01,
            bytes32(uint256(1)),
            bytes32(uint256(2))
        );

        (bool exists, bytes32 categoryBitMap) = registry.tryGetCategoryBitmap(
            0x01,
            bytes32(uint256(1))
        );
        assertTrue(exists == true);
        assertTrue(categoryBitMap == bytes32(uint256(2)));
    }
}
