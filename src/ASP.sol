// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./RecordCategoryRegistry.sol";

contract ASP is RecordCategoryRegistry {
    constructor() RecordCategoryRegistry(msg.sender) {}

    /*//////////////////////////////////////////////////////////////////////////
                     ***ASSOCIATION-SET RELEVANT PUBLIC FUNCTIONS***
    //////////////////////////////////////////////////////////////////////////*/
    /**
     * @notice Applies a predicate to filter elements based on their properties
     *
     * @dev Uses bitwise operations to perform set operations on element properties
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
    ) external view returns (bytes32[] memory elements, uint256 setCardinality) {
        bytes32[] memory satisfyingElements = new bytes32[](subset.length);
        setCardinality = 0;

        for (uint256 i = 0; i < subset.length; i++) {
            bytes32 element = subset[i];
            (bool isMember, bytes32 elementProperties) = getCategoryForRecordHash(domain, element);
            if (!isMember) {
                continue;
            }
            if (_applyPredicate(predicateType, characteristicFunction, elementProperties)) {
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
                                INTERNAL FUNCTIONS
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
    function _applyPredicate(PredicateType predicateType, bytes32 characteristicFunction, bytes32 elementProperties)
        internal
        pure
        returns (bool satisfiesPredicate)
    {
        if (predicateType == PredicateType.Intersection) {
            // Intersection: element must have all properties defined in the characteristic function
            satisfiesPredicate = (elementProperties & characteristicFunction) == characteristicFunction;
        } else if (predicateType == PredicateType.Union) {
            // Union: element must have at least one property defined in the characteristic function
            satisfiesPredicate = (elementProperties & characteristicFunction) != 0;
        } else if (predicateType == PredicateType.Complement) {
            // Complement: element must not have any properties defined in the characteristic function
            satisfiesPredicate = (elementProperties & characteristicFunction) == 0;
        }
        return false;
    }
}
