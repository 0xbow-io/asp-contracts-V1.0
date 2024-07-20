# ASP Store V1.0

> This repository contains the solidity contracts utilised by the Association
> Set Prorivder (ASP) to support Protocol Integrations.
> It also contains go & ts packages that provides the necessary
> interfaces to interact with the contracts.

## Context

The ASP defines a protocol as a class of functional state machine with a unique
identifier called "scope".

> A Protocol instances are on-chain implementations of this state machine
> class.
> For example: Uniswap V2 is a protocol, and a UniV2 Pool in mainnet is an
> instance of this protocol.
> A scope value can be computed for that instance as the Keccak256 Hash of
> (chainID, contract address, TokenA address, TokenB Address)

The key attributes of the protocol:

- scope: Unique identifier for each protocol instance (e.g. Keccak256(chainID,
  address) % SNARK_SCALAR_FIELD)
- stateSize: Measurable size of instance state
- stateRoot: Unique reference to current state

Agents serve as the input sources to invoke state transitions.

> [!IMPORTANT]

> **Protocols must explicitly define & announce state-transitions onchain via emitting "Record" events.**

These events contains:

- Request: Reference to processed input
  - Note that for zK-based protocols, the complete set of input values cannot be included in the Record event.
  - Hence a generic "Request" reference is used instead.
- New stateRoot
- New stateSize

The "association set-provider" (ASP) is an external system:

- Which is aware of all Record Events emitted by protocol instances.
  - A set of hashes of all Record Events which is
- And applies predicate function to catogrise Record Events.
- Record categories are stored on-chain and is publicaly accessible.
- Protocols can utilise the ASP to support their implementation of a zero-knowledge proofs of association (zk-PoA) for agents.
  - **zk-PoA enables privacy-preserving protocol to protocol communication**
  - i.e.: Privacy Pool zk-PoA implementation.

---

# Guide to Protocol x ASP Integration

#### **Prerequisites**

Before attempting any integration, these prerequisites must be met:

- [x] The protocol behaves as a functional state machine with clearly defined inputs & state-transitions (i.e. Liquidity Pools, Vaults, DAO's).
- [x] A unique Scope identifier can be computed for the protocol instances.
- [x] At every **state-transition** a compatible "Record" event is emitted on-chain either by the instance itself or a proxy.
- [x] The value of stateRoot & stateSize in the "Record" event corectly reflect the instance state.
- [x] A unique hash can be computed from the "Record" event

#### **Initialisation**

Request ASP to add a new store for the protocol, communicating the following:

- The scope value for the protocol instance.
- The address of the contract that emits the Record event
- Reference to the record event meta (or the meta itself). The meta is information on the following:
  - The event signature (e.g. `event Record(uint256 scope, bytes32 request, bytes32 stateRoot, uint256 stateSize);`)
  - The event hash function (e.g. `keccak256(abi.encodePacked(scope, request, stateRoot, stateSize)) or PosiedonHash...`)
  - The predicate function that categorises the record event
  - All the possible categories that the record event can be classified into

The ASP should review & approve the request and create a new store for the protocol.

> **This is announced by the emisison of the NewStore event**

The ASP then will kick-off the relevant services to aggregate all the record events, categorise them and store the categories into the ASP Store contract.

> **This is announced by the emisison of the NewRecord event**

#### **Utility**

- **Querying**: The ASP provides a set of public functions to query the categories of the record events.
- **Inclusion Proof**: An inclusion proof can be generated off-chain to verify that a record event is known to the ASP

Here are some case studies of protocol integrations:

#### Uniswap V2 Protocol

### ERC-4337

### Case Study: Privacy Pool Integration
