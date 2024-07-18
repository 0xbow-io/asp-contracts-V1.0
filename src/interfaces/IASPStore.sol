// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @title IASPStore contract interface.
interface IASPStore {
    error SenderIsNotAdmin(address sender);
    error InsufficientRole(address sender, bytes32 role);
    error ScopeStoreExists(uint256 scope);
    /// @dev NewStore event is the announcement of
    /// of a new store for a particular protocol instance
    /// referenced by its scope
    /// @param scope scope of the protocol's instance
    /// @param recordEventEmitter address of the conttract that emits the record events
    /// @param meta meta data of the instance
    /// @notice this event is emitted when a new store is created
    /// for a particular protocol instance
    /// @notice the scope is a unique identifier for the protocol instance
    /// @notice the instance is the address of the protocol instance
    /// @notice the meta is the meta data of the protocol instance
    /// - this could be a IPFS uri / hash
    /// - or a JSON string containing the protocol's metadata
    event NewStore(
        uint256 scope,
        address recordEventEmitter,
        string meta,
        uint256 storeID
    );

    /// @dev NewRecord event is the announcement of the storage
    /// of a new record hash for a particular protocol instance
    event NewRecord(
        uint256 scope, /// scope of the protocol's instance
        uint256 recordSetRoot, /// updated root of the merkle tree from set of record Hashes
        uint256 recordSetSize, /// size of the record set
        bytes32 hash, /// hash of the record
        bytes32 category, /// category of the record
        bytes32 srcTxHash, /// reference to the original Tx hash which contains this record event
        bytes32 srcLogIndex /// reference to the original log index which contains this record event
    );

    /// @dev UpdatedRecord announces when the category of the record hash is updated
    event UpdatedRecord(
        uint256 scope, /// scope of the protocol's instance
        bytes32 hash, /// hash of the record
        bytes32 category /// category of the record
    );
}
