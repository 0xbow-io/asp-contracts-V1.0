// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/*  OpenZeppelin Imports */
import {EnumerableMap} from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import {AccessControlDefaultAdminRules} from "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
/*  Zk-kit Imports */
import {InternalLeanIMT, LeanIMTData} from "zk-kit.solidity/packages/lean-imt/contracts/InternalLeanIMT.sol";

/**
 * @title   RecordCategoryRegistry contract
 * @notice  This contract can be deployed by any ASP to support
 *          onchain protocol integrations.
 * @notice  ⚠️ This contract is not audited and should be used with caution.
 * @dev This contract stores the categories for all known records
 *      hashes for a given scope.
 *      Set of record hashes within a scope is used to ccnstruct a
 *      lean-IMT (Incremental Merkle Tree).
 * @dev ⚠️ This contract implements roles to control certain actions
 *      such as creating stores and inserting / updating new records.
 *      Default admin role is set to the deployer of the contract.
 *      Admin also has the ability to grant the store adnmin role.
 *
 * @author  oxbow.io (https://github.com/0xbow-io/asp-store-v1)
 */
contract RecordCategoryRegistry is AccessControlDefaultAdminRules {
    using EnumerableMap for EnumerableMap.Bytes32ToBytes32Map;
    using InternalLeanIMT for LeanIMTData;

    error SenderIsNotAdmin(address sender);
    error InsufficientRole(address sender, bytes32 role);
    error RecordNotFound(uint256 scope, bytes32 recordHash);

    /// @notice Emitted when a record category is set or updated
    /// @param scope The protocol scope identifier
    /// @param root The new Merkle root after the update
    /// @param totalRecords The total number of records after the
    event Update(uint256 scope, uint256 root, uint256 totalRecords);

    /// @dev role identifier used to grant addresses ability to
    ///      update records in the store & recordTrees,
    bytes32 public constant REGISTRY_ADMIN_ROLE =
        keccak256("REGISTRY_ADMIN_ROLE");

    /// @notice Mapping of protocol scopes to their record categories
    /// @dev This mapping stores the category bitmaps for each record hash within a protocol scope
    /// @dev The outer mapping key is the protocol scope (uint256)
    /// @dev The inner mapping uses EnumerableMap.Bytes32ToBytes32Map xor efficient enumeration
    /// @dev Record hashes (bytes32) are mapped to their category bitmaps (bytes32)
    /// @dev Each bit in the category bitmap represents a specific category that is defined by ASP / protocol
    /// @dev Record can have multiple categories, which are represented by setting multiple bits in the bitmap
    /// @dev Using EnumerableMap library for easy iteration over the set of record hashes
    /// @dev For example:
    ///      - bit 0 (000..001): fee paid
    ///      - bit 1 (000..010): approved
    ///      - bit 2 (000..100): processed
    ///      - bit 3 (000..101): processed and fee paid
    /// @notice  ⚠️ records cannot be removed or deleted.
    mapping(uint256 scope => EnumerableMap.Bytes32ToBytes32Map recCatMap)
        private scopeRecordCategories;

    /// @notice Mapping of protocol scopes to their Merkle trees of record hashes
    /// @dev This mapping stores Merkle trees for efficient proof generation of record inclusion
    /// @dev The key is the protocol scope (uint256)
    /// @dev The value is a LeanIMTData struct, representing a Lean Incremental Merkle Tree
    /// @dev Each leaf in the Merkle tree corresponds to a record hash
    /// @dev This structure allows for:
    ///      - Efficient updates when new records are added
    ///      - Generation of inclusion proofs for any record
    ///      - Verification that a record belongs to a specific protocol scope
    /// @dev ⚠️ The Lean IMT library does not provide a straightfoward method to
    ///     iterate through the tree leaf nodes.
    ///     Instead iterate through recCatMap will iterate through the leaf nodes.
    ///     using the at(..) function from EnumerableMap.Bytes32ToBytes32Map library.
    mapping(uint256 scope => LeanIMTData merkleTree)
        private scopeRecordMerkleTrees;

    /*//////////////////////////////////////////////////////////////////////////
    | MODIFIERS
    //////////////////////////////////////////////////////////////////////////*/

    /// @dev modifier to check that the caller has the admin role
    modifier adminOnly() {
        if (!hasRole(DEFAULT_ADMIN_ROLE, msg.sender)) {
            revert SenderIsNotAdmin(msg.sender);
        }
        _;
    }

    /// @dev modifier to check that the caller has a specific role
    /// @param role is the role to check for
    modifier sufficientRole(bytes32 role) {
        if (!hasRole(role, msg.sender)) {
            revert InsufficientRole(msg.sender, role);
        }
        _;
    }

    /*//////////////////////////////////////////////////////////////////////////
    | CONTRACT CONSTRUCTOR
    //////////////////////////////////////////////////////////////////////////*/

    constructor(
        address registryAdmin
    )
        /// Explicit initial `DEFAULT_ADMIN_ROLE` holder
        AccessControlDefaultAdminRules(3 days, msg.sender)
    {
        /// Grant the contract deployer the ability to crhasRoleeate stores
        grantRegistryAdminRole(registryAdmin);
    }

    /*//////////////////////////////////////////////////////////////////////////
    | ADMIN FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/

    /// @dev GrantRegistryAdminRole grants the REGISTRY_ADMIN_ROLE to an address
    /// @param account the address to grant the role to
    function grantRegistryAdminRole(address account) public adminOnly {
        grantRole(REGISTRY_ADMIN_ROLE, account);
    }

    /// @dev RevokeRegistryAdminRole revokes the REGISTRY_ADMIN_ROLE from an address
    /// @param account the address to revoke the role from
    function revokeRegistryAdminRole(address account) public adminOnly {
        revokeRole(REGISTRY_ADMIN_ROLE, account);
    }

    /*//////////////////////////////////////////////////////////////////////////
    | REGISTRY_ADMIN_ROLE FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/

    /**
     * @notice Sets or updates the category for a record within a specific protocol scope
     *
     * @dev This function adds a new record or updates an existing one's category, and
     *     updates the Merkle tree if necessary.
     *
     * @param scope The unique identifier of the protocol instance
     * @param recordHash The hash of the record event
     * @param categoryBitmap A bytes32 representing the categories as a bitmap
     * @custom:access ⚠️ Restricted to accounts with REGISTRY_ADMIN_ROLE
     *
     */
    function setCategoryForRecord(
        uint256 scope,
        bytes32 recordHash,
        bytes32 categoryBitmap
    ) public sufficientRole(REGISTRY_ADMIN_ROLE) {
        // get current merkle root
        uint256 root = scopeRecordMerkleTrees[scope]._root();
        /// .set() will return true if recordHash is not existing key
        /// if not an existing key, it means the record hash is new
        if (scopeRecordCategories[scope].set(recordHash, categoryBitmap)) {
            /// update the merkle tree with the new record hash
            /// and apply the new root value
            root = scopeRecordMerkleTrees[scope]._insert(uint256(recordHash));
        }
        /// announce update by emitting an Update event
        emit Update(scope, root, scopeRecordCategories[scope].length());
    }

    /*//////////////////////////////////////////////////////////////////////////
    | PUBLIC FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/

    /**
     * @notice Returns the category bitmap for a record hash for a specific protocol scope
     * @param scope The protocol scope identifier
     * @param recordHash The hash of the record event
     * @return categoryBitmap The category bitmap for the record hash
     */
    function getCategoryBitmap(
        uint256 scope,
        bytes32 recordHash
    ) public view returns (bytes32 categoryBitmap) {
        (bool exists, bytes32 bitmap) = scopeRecordCategories[scope].tryGet(
            recordHash
        );
        if (!exists) {
            revert RecordNotFound(scope, recordHash);
        }
        return bitmap;
    }

    /**
     * @notice Returns the category bitmap for a record hash at a given index
     *          for a specific protocol scope
     * @param scope The protocol scope identifier
     * @param index The index of the record hash
     * @return recordHash recordHash at the given index
     * @return categoryBitmap The category bitmap for the record hash
     */
    function getRecordHashAndCategoryAt(
        uint256 scope,
        uint256 index
    ) public view returns (bytes32 recordHash, bytes32 categoryBitmap) {
        return scopeRecordCategories[scope].at(index);
    }

    /**
     * @notice Return the category bitmap for a record hash
     *          for a specific protocol scope
     # @dev does not revert if the record hash does not exist
     * @param scope The protocol scope identifier
     * @param recordHash The hash of the record event
     * @return exists A boolean indicating if the record hash exists
     * @return categoryBitmap The category bitmap for the record hash
     */
    function tryGetCategoryBitmap(
        uint256 scope,
        bytes32 recordHash
    ) public view returns (bool exists, bytes32 categoryBitmap) {
        return scopeRecordCategories[scope].tryGet(recordHash);
    }

    /**
     * @notice Returns the record hashes and their categories for a specific protocol scope between
     *         a given range
     * @param scope The protocol scope identifier
     * @param from The start index of the range
     * @param to The end index of the range
     * @return recordHashes The record hashes for the given range
     * @return categoryBitmaps The category bitmaps for the given range
     */
    function getRecordHashesAndCategories(
        uint256 scope,
        uint256 from,
        uint256 to
    )
        public
        view
        returns (
            bytes32[] memory recordHashes,
            bytes32[] memory categoryBitmaps
        )
    {
        require(
            from < to && to <= scopeRecordCategories[scope].length(),
            "Invalid range"
        );
        uint256 length = to - from;
        recordHashes = new bytes32[](length);
        categoryBitmaps = new bytes32[](length);

        for (uint256 i = 0; i < length; i++) {
            (recordHashes[i], categoryBitmaps[i]) = scopeRecordCategories[scope]
                .at(from + i);
        }
    }

    /**
     * @notice Returns the last record hash & it's category and the merkle-root
     *         for a specific protocol scope
     * @param scope The protocol scope identifier
     * @return root The merkle root for the protocol scope
     * @return recordHash The hash of the last known record event
     * @return categoryBitmap The category bitmap for the last known record event
     * @return index The index of the last known record event
     */
    function getLatestForScope(
        uint256 scope
    )
        public
        view
        returns (
            uint256 root,
            bytes32 recordHash,
            bytes32 categoryBitmap,
            uint256 index
        )
    {
        root = scopeRecordMerkleTrees[scope]._root();
        index = scopeRecordCategories[scope].length();
        if (index > 0) {
            (recordHash, categoryBitmap) = scopeRecordCategories[scope].at(
                index - 1
            );
        }
    }
}
