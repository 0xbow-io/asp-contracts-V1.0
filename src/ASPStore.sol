// SPDX-License-Identifier: MITs
pragma solidity ^0.8.4;

/*  OpenZeppelin Imports */
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {EnumerableMap} from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import {AccessControlDefaultAdminRules} from "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
/*  Zk-kit Imports */
import {InternalLeanIMT, LeanIMTData} from "zk-kit.solidity/packages/lean-imt/contracts/InternalLeanIMT.sol";
import "./interfaces/IASPStore.sol";

/**
 * @title ASPStore contract.
 * @notice This contract is responsible for storing the Protocol relevant data.
 * such as inclusion / exclusion sets, etc.
 * @dev This contract is a core component of the ASP x Protocol integration.
 **/
contract ASPStore is AccessControlDefaultAdminRules, IASPStore {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableMap for EnumerableMap.Bytes32ToBytes32Map;
    using EnumerableMap for EnumerableMap.UintToUintMap;
    using InternalLeanIMT for LeanIMTData;

    /// role for the store creator
    bytes32 public constant STORE_CREATOR_ROLE =
        keccak256("STORE_CREATOR_ROLE");

    /// @title RecordStore struct
    /// @dev RecordStore is a struct that holds
    /// - the record hash to category mapping
    /// - set of record hashes represented as a merkle tree
    /// * [WARNING]
    /// * Records should not be removed from recordToCategory or the recordTree.
    /// * leafs of the recordTree should not be updated, only new records should be added.
    /// * assumed that the ordering of the the record hahses follows the ordering of the
    /// * tx hash & log index of the record event emission.
    /// * this is not checked in the contract.
    struct RecordStore {
        /// @dev recordTree is a merkle tree of record hashes
        LeanIMTData recordTree;
        /// @dev: mapping of Record Hash to a category
        /// using bytes32 to provide flexibility for the category (i.e. for bit-map)
        /// category is defined by the protocol & ASP.
        /// using EnumerableMap lib for easier iteration over record hashes
        EnumerableMap.Bytes32ToBytes32Map recordToCategory;
    }

    /// @dev stores stores the RecordStore for each known scope
    RecordStore[] stores;
    /// @dev scopeToStoreIdx is a mapping of scope to store index
    EnumerableMap.UintToUintMap scopeToStoreIdx;

    constructor()
        AccessControlDefaultAdminRules(
            3 days,
            msg.sender // Explicit initial `DEFAULT_ADMIN_ROLE` holder
        )
    {
        // Grant the contract deployer the ability to create stores
        GrantStoreCreatorRole(msg.sender);
    }

    /*//////////////////////////////////////////////////////////////////////////
                            MODIFIERS
    //////////////////////////////////////////////////////////////////////////*/
    modifier AdminOnly() {
        if (!hasRole(DEFAULT_ADMIN_ROLE, msg.sender)) {
            revert SenderIsNotAdmin(msg.sender);
        }
        _;
    }

    modifier SufficientRole(bytes32 role) {
        if (!hasRole(role, msg.sender)) {
            revert InsufficientRole(msg.sender, role);
        }
        _;
    }

    /*//////////////////////////////////////////////////////////////////////////
                            PUBLIC FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/
    function GrantStoreCreatorRole(address acocunt) public AdminOnly {
        grantRole(STORE_CREATOR_ROLE, acocunt);
    }

    /**
     * @dev CreateStore creates a new store for a particular protocol instance
     * @param scope scope of the protocol's instance
     * @param recordEventEmitter address of the conttract that emits the record events
     * @param meta meta data of the instance
     * @return storeID the ID of the store
     * [WARNING]
     * only addresses with the STORE_CREATOR_ROLE can create a store
     * by default this is the contract deployer (or the ASP entity)
     **/
    function CreateStore(
        uint256 scope,
        address recordEventEmitter,
        string calldata meta
    ) public SufficientRole(STORE_CREATOR_ROLE) returns (uint256 storeID) {
        require(hasRole(STORE_CREATOR_ROLE, msg.sender));
        storeID = stores.length;
        stores.push();
        if (!scopeToStoreIdx.set(scope, storeID)) {
            revert ScopeStoreExists(scope);
        }
        emit NewStore(scope, recordEventEmitter, meta, storeID);
    }

    /**
     * @dev InsertNewRecord creates a new store for a particular protocol instance
     * @param storeID the ID of the store
     * @param recordHash the hash of the record
     * @param category the category of the record
     **/
    function InsertNewRecord(
        uint256 storeID,
        bytes32 recordHash,
        bytes32 category
    ) public SufficientRole(STORE_CREATOR_ROLE) {
        require(hasRole(STORE_CREATOR_ROLE, msg.sender));
    }

    /*//////////////////////////////////////////////////////////////////////////
                            INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/
}
