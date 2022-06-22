package block

import (
	"fmt"
	"reflect"

	"github.com/cybriq/kismet/pkg/blockinterface"
	"github.com/cybriq/kismet/pkg/bytes"
	"github.com/cybriq/kismet/pkg/ed25519"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/known"
	"github.com/cybriq/kismet/pkg/proof"
	"lukechampine.com/blake3"
)

// Block is the base block structure, which can be extended for specific types
type Block struct {
	// Type is the type code
	Type known.Type

	// Time is a unix 64 bit timestamp in nanoseconds that can measure until 2262 AD.
	Time int64

	// Difficulty is the difficulty target created by the previous Block(s)
	// Difficulty multiplied by the divergence from the time target using a
	// Proportional/Integral formula to derive the valid difficulty on a subsequent
	// block
	Difficulty hash.Hash

	// Previous is the previous block(s) this block is mined on, and for Proposal
	// and Congress tokens, the last element is the IPFS hash of the proposal
	Previous hash.Hash

	// PublicKey is the identity that can sign for this block
	ed25519.PublicKey
}

var _ blockinterface.Blocker = &Block{}

var Name = reflect.TypeOf(Block{}).Name()

const Length = 2 + 8 + hash.Len + hash.Len + ed25519.PublicKeySize

func (b *Block) Marshal() (by []byte, err error) {

	if b == nil {
		// It is programmer error if a nil pointer is passed
		err = fmt.Errorf("cannot marshal nil Block")
		log.E.Ln(err)
		return
	}

	by = make([]byte, b.Length())

	by[0] = byte(b.Type)
	by[1] = byte(b.Type >> 8)
	copy(by[2:10], bytes.FromInt64(b.Time))

	// The rest are simple copy operations
	copy(by[10:hash.Len+10], b.Difficulty[:])
	copy(by[42:hash.Len+42], b.Previous[:])
	copy(by[74:ed25519.PublicKeySize+74], b.PublicKey[:])
	return

}

func (b *Block) Unmarshal(by []byte) (err error) {

	if len(by) != b.Length() {
		err = fmt.Errorf(
			"data length incorrect, got %d expected %d",
			len(by), b.Length(),
		)
		log.E.Ln(err)
		return
	}

	*b = Block{}

	// again, just doing this directly is the fastest.
	b.Type = known.Type(by[0]) + known.Type(by[1])<<8
	b.Time = bytes.ToInt64(by[2:10])

	// The hashes and keys are just copy operations
	copy(b.Difficulty[:], by[10:10+hash.Len])
	copy(b.Previous[:], by[42:hash.Len+42])
	copy(b.PublicKey[:], by[74:ed25519.PublicKeySize+74])

	return
}

func (b *Block) Length() (l int) { return Length }

func (b *Block) ID() string { return Name }

// PoWHash returns the Proof of Work hash for a given block. We would not use
// this function in a miner because only the timestamp and previous needs to be
// changed between attempts, and would be faster to directly change the bytes.
func (b *Block) PoWHash() (h []byte, err error) {

	var by []byte
	if by, err = b.Marshal(); log.E.Chk(err) {
		return
	}

	h = proof.DivHash4(by[:])
	return
}

// IndexHash is the hash that is used to store and search for blocks, as it is
// far faster to calculate than the PoWHash.
func (b *Block) IndexHash() (h hash.Hash, err error) {

	var by []byte
	if by, err = b.Marshal(); log.E.Chk(err) {

		return
	}

	h = blake3.Sum256(by[:])
	return
}

// GetBlock does nothing interesting in this implementation, but in derived,
// extended block formats with extra fields, this would return the embedded
// Block that is extended from.
func (b *Block) GetBlock() *Block { return b }
