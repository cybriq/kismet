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

// Hash is a 32 byte hash.
type Hash [32]byte

// Block is the base block structure, which can be extended for specific types
type Block struct {
	// Time is a unix 64 bit timestamp in nanoseconds that can measure until
	// 2262 AD.
	Time int64

	// Type is the type code - constants ending in Block, above
	Type byte

	// Count is the number of blocks this block refers back to
	Count byte

	// Difficulty is the difficulty target created by the previous Block(s)
	// Difficulty multiplied by the divergence from the time target using a
	// Proportional/Integral formula
	Difficulty Hash

	// Previous is the previous block(s) this block is mined on, and for Proposal and
	// Congress tokens, the last element is the IPFS hash of the proposal
	Previous []Hash

	// PublicKey is in fact always 256 bits/32 bytes long but
	// github.com/cloudflare/circl implementation does not use an array
	ed25519.PublicKey
}
