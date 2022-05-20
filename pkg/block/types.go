package block

import (
    "github.com/cloudflare/circl/sign/ed25519"
)

// The following values signify block types
const (
    // ValidatorBlock is for a slot in the block production schedule
    ValidatorBlock = 1

    // ProposalBlock is for a governance proposal
    ProposalBlock = 2

    // CongressBlock is for a voting right on a given submitted proposal
    CongressBlock = 3
)

// MergedBlock means that this block merges a fork
const MergedBlock = 128

// Hash is a 32 byte hash.
type Hash [32]byte

// Block is the base block structure, which can be extended for specific types
type Block struct {

    // Previous is the primary previous block this block is mined on
    Previous Hash

    // Type is the type code - constants ending in Block, above
    Type byte

    // Time is a unix 64 bit timestamp in nanoseconds that can measure until
    // 2262 AD. Time of blocks must always progress to be valid
    Time int64

    // Difficulty is the difficulty target created by the previous Block(s)
    // target multiplied by the divergence from the time target
    Difficulty Hash

    // PublicKey is in fact always 256 bits/32 bytes long but
    // github.com/cloudflare/circl implementation does not use an array
    ed25519.PublicKey
}

// Merge is a block that joins two blocks that have formed a fork. These blocks
// have double difficulty compared to the average of the two parent blocks, and
// supersede single parent blocks when there is a fork.
type Merge struct {
    Block
    SecondPrevious Hash
}

// Congress blocks issue the right to a vote against a specified Proposal. These
// are only valid if the proposal exists on the pBFT ledger.
type Congress struct {
    Block
    Proposal Hash
}

// CongressMerge is a Congress block that merges a fork, it will supersede
// extensions when a fork exists. Difficulty is double the average of the parent
// blocks.
type CongressMerge struct {
    Congress
    SecondPrevious Hash
}
