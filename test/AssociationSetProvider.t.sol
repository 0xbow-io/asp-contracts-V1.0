// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/AssociationSetProvider.sol";
import "forge-std/console.sol";

contract TestAssociationSetProvider is Test {
    AssociationSetProvider asp;
    uint256 constant DOMAIN = 1;

    function setUp() public {
        asp = new AssociationSetProvider();
        // Setup some test data
        asp.setCategoryForRecord(
            DOMAIN,
            bytes32(uint256(1)),
            bytes32(uint256(0x1))
        ); // Element 1: Category 1 0b001
        asp.setCategoryForRecord(
            DOMAIN,
            bytes32(uint256(2)),
            bytes32(uint256(0x2))
        ); // Element 2: Category 2 0b010
        asp.setCategoryForRecord(
            DOMAIN,
            bytes32(uint256(3)),
            bytes32(uint256(0x3))
        ); // Element 3: Categories 1 and 2 0b011
        asp.setCategoryForRecord(
            DOMAIN,
            bytes32(uint256(4)),
            bytes32(uint256(0x4))
        ); // Element 4: Category 3 0b100
        asp.setCategoryForRecord(
            DOMAIN,
            bytes32(uint256(5)),
            bytes32(uint256(0x7))
        ); // Element 5: Categories 1, 2, and 3 0b111
    }

    function test_IntersectionPositive() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }

        (bytes32[] memory result, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x3)), // Looking for elements with categories 1 AND 2 0b011
            AssociationSetProvider.PredicateType.Intersection
        );

        assertEq(count, 2);
        assertEq(uint256(result[0]), 3);
        assertEq(uint256(result[1]), 5);
    }

    function test_IntersectionNegative() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }

        (bytes32[] memory result, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x5)), // Looking for elements with categories 1 AND 3 0b101
            AssociationSetProvider.PredicateType.Intersection
        );

        assertEq(count, 1);
        assertEq(uint256(result[0]), 5);
    }

    function test_UnionPositive() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }

        (bytes32[] memory result, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x3)), // Looking for elements with categories 1 OR 2
            AssociationSetProvider.PredicateType.Union
        );

        assertEq(count, 4);
        assertEq(uint256(result[0]), 1);
        assertEq(uint256(result[1]), 2);
        assertEq(uint256(result[2]), 3);
        assertEq(uint256(result[3]), 5);
    }

    function test_UnionNegative() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }
        (, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x8)), // Looking for elements with category 4 (which doesn't exist)
            AssociationSetProvider.PredicateType.Union
        );

        assertEq(count, 0);
    }

    function test_ComplementPositive() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }

        (bytes32[] memory result, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x4)), // Looking for elements without category 3
            AssociationSetProvider.PredicateType.Complement
        );

        assertEq(count, 3);
        assertEq(uint256(result[0]), 1);
        assertEq(uint256(result[1]), 2);
        assertEq(uint256(result[2]), 3);
    }

    function test_ComplementNegative() public {
        bytes32[] memory subset = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            subset[i] = bytes32(uint256(i + 1));
        }

        (, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0)), // Looking for elements without any categories (which doesn't exist)
            AssociationSetProvider.PredicateType.Complement
        );

        assertEq(count, 5);
    }

    function test_EmptySubset() public {
        bytes32[] memory subset = new bytes32[](0);

        (, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x1)),
            AssociationSetProvider.PredicateType.Union
        );

        assertEq(count, 0);
    }

    function test_NonExistentElement() public {
        bytes32[] memory subset = new bytes32[](1);
        subset[0] = bytes32(uint256(100)); // This element doesn't exist in our setup

        (, uint256 count) = asp.applyPredicate(
            DOMAIN,
            subset,
            bytes32(uint256(0x1)),
            AssociationSetProvider.PredicateType.Union
        );

        assertEq(count, 0);
    }
}
