# ASP Store

> This repository contains the solidity contracts utilised by the Association
Set Provider to store Inclusion / exclusion sets.
> It also contains go & ts packages that provides the necessary
interfaces to interact with the contracts.

## Context

A protocol is defined as a class of functional state machine with a unique
identifier called "scope".

> Protocol instances are on-chain implementations of this state machine class.

The key attributes of the protocol:
* scope: Unique identifier for each protocol instance (e.g. Keccak256(chainID,
address) % SNARK_SCALAR_FIELD)
* stateSize: Measurable size of instance state
* stateRoot: Unique reference to current state

Agents serve as the input sources to invoke state transitions.

> [!IMPORTANT]

> **Protocols must explicitly define & log state-transitions onchain via emitting "Record" events.**

These events contains:
* Request: Reference to processed input
  * Note that for zK-based protocols, the complete set of input values cannot be included in the Record event.
  * Hence the "Request" reference is used instead.
* New stateRoot
* New stateSize

The "association set-provider" (ASP) is an external system:
* That applies a predicate function to generate an inclusion set of Record events
* And represents inclusion set as a Merkle tree of event hashes

With the ASP, agents can generate zero-knowledge proofs of association (zk-PoA) which proves:
* Ownership of inputs that invoked specific state transitions
* Sufficient membership of Record events (associated with those state transitions) in the ASP inclusion set

__zk-PoA enables privacy-preserving integration between protocol instances to
support transmission & storage of confidential data__
