// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/* Local Imports */
import "./RecordCategoryRegistry.sol";
import "forge-std/console.sol";

/**
 * @title   AssociationSetProvider contract
 * @notice  This contract can be deployed by any ASP to support
 *          onchain protocol integrations.
 * @notice  ⚠️ This contract is not audited and should be used with caution.
 * @dev This contract stores the categories for all known records
 *      hashes for a given scope.
 *      Set of record hashes within a scope is used to ccnstruct a
 *      lean-IMT (Incremental Merkle Tree).
 * @dev ⚠️ This contract inherits RecordCategoryRegistry which implements roles
 *      to control certain actions such as creating stores and
 *      inserting / updating new records.
 *      Default admin role is set to the deployer of the contract.
 *      Admin also has the ability to grant the registry adnmin role.
 *
 * @author  oxbow.io (https://github.com/0xbow-io/asp-contract-V1.0)
 */
contract AssociationSetProvider is RecordCategoryRegistry {
    /// @notice Defines the type of predicate to apply
    /// @dev Intersection: all bits must match
    /// @dev Union: at least one bit must match,
    /// @dev Complement: no bits should match
    enum PredicateType {
        Intersection,
        Union,
        Complement
    }

    constructor() RecordCategoryRegistry(msg.sender) {}

    /*//////////////////////////////////////////////////////////////////////////
    | PUBLIC FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/
    /**
     * @notice Applies a predicate to filter elements based on their properties
     *
     * @dev Uses bitwise operations to perform set operations on element properties.
     *      Accesses the RecordCategoryRegistry to get the properties of each element.
     *
     * @param domain The context or scope in which to apply the predicate
     * @param subset Array of elements to filter
     * @param characteristicFunction Bitmap representing the properties to filter by
     * @param predicateType The type of set operation to perform (Intersection, Union, or Complement)
     * @return elements The elements that satisfy the predicate
     * @return setCardinality The number of elements that satisfy the predicate
     *
     */
    function applyPredicate(
        uint256 domain,
        bytes32[] calldata subset,
        bytes32 characteristicFunction,
        PredicateType predicateType
    ) public view returns (bytes32[] memory elements, uint256 setCardinality) {
        bytes32[] memory satisfyingElements = new bytes32[](subset.length);
        for (uint256 i = 0; i < subset.length; i++) {
            bytes32 element = subset[i];
            (bool isMember, bytes32 elementProperties) = tryGetCategoryBitmap(
                domain,
                element
            );
            if (!isMember) {
                continue;
            }
            if (
                _applyPredicate(
                    predicateType,
                    characteristicFunction,
                    elementProperties
                )
            ) {
                satisfyingElements[setCardinality] = element;
                setCardinality++;
            }
        }
        // Resize the array to remove empty slots (maintain proper cardinality)
        assembly {
            mstore(satisfyingElements, setCardinality)
        }

        return (satisfyingElements, setCardinality);
    }

    /*//////////////////////////////////////////////////////////////////////////
    | INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/
    /**
     * @notice Checks if an element satisfies a predicate based on its properties
     *
     * @dev Uses bitwise operations to perform set operations on element properties
     *
     * @param characteristicFunction Bitmap representing the properties to filter by
     * @param predicateType The type of set operation to perform (Intersection, Union, or Complement)
     * @param elementProperties Bitmap representing the properties of the element
     * @return satisfiesPredicate Whether the element satisfies the predicate
     *
     */
    function _applyPredicate(
        PredicateType predicateType,
        bytes32 characteristicFunction,
        bytes32 elementProperties
    ) internal pure returns (bool satisfiesPredicate) {
        if (predicateType == PredicateType.Intersection) {
            // Intersection: element must have all properties defined in the characteristic function
            satisfiesPredicate =
                (elementProperties & characteristicFunction) ==
                characteristicFunction;
        } else if (predicateType == PredicateType.Union) {
            // Union: element must have at least one property defined in the characteristic function
            satisfiesPredicate =
                (elementProperties & characteristicFunction) != 0;
        } else if (predicateType == PredicateType.Complement) {
            // Complement: element must not have any properties defined in the charact
            // eristic function
            satisfiesPredicate =
                (elementProperties & characteristicFunction) == 0;
        }
    }
}
