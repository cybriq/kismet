package block

import (
	"fmt"
	"unsafe"

	"github.com/cybriq/kismet/pkg/blockinterface"
	"github.com/cybriq/kismet/pkg/ed25519"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/known"
	"github.com/cybriq/kismet/pkg/proof"
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

	// PublicKey is in fact always 256 bits/32 bytes long but
	// github.com/cloudflare/circl implementation does not use an array
	ed25519.PublicKey
}

var _ block.Blocker = &Block{}

const Name = "kismet.Block"

func (b *Block) Marshal() (bytes []byte, err error) {

	if b == nil {
		// It is programmer error if a nil pointer is passed
		err = fmt.Errorf("cannot marshal nil Block")
		log.E.Ln(err)
		return
	}

	bytes = make([]byte, b.Length())

	// There is functions to do these but that would be slower than doing this
	// directly with the integers
	bytes[0] = byte(b.Type)
	bytes[1] = byte(b.Type >> 8)
	bytes[2] = byte(b.Time)
	bytes[3] = byte(b.Time >> 8)
	bytes[4] = byte(b.Time >> 16)
	bytes[5] = byte(b.Time >> 24)
	bytes[6] = byte(b.Time >> 32)
	bytes[7] = byte(b.Time >> 40)
	bytes[8] = byte(b.Time >> 48)
	bytes[9] = byte(b.Time >> 56)

	// The rest are simple copy operations
	copy(bytes[10:hash.Len+10], b.Difficulty[:])
	copy(bytes[42:hash.Len+42], b.Previous[:])
	copy(bytes[74:ed25519.PublicKeySize+74], b.PublicKey[:])
	return

}

func (b *Block) Unmarshal(bytes []byte) (err error) {

	if len(bytes) != b.Length() {
		err = fmt.Errorf(
			"data length incorrect, got %d expected %d",
			len(bytes), b.Length(),
		)
		log.E.Ln(err)
		return
	}

	*b = Block{}

	// again, just doing this directly is the fastest.
	b.Type = known.Type(bytes[0]) + known.Type(bytes[1])<<8
	b.Time = int64(bytes[2]) +
		int64(bytes[3])<<8 +
		int64(bytes[4])<<16 +
		int64(bytes[5])<<24 +
		int64(bytes[6])<<32 +
		int64(bytes[7])<<40 +
		int64(bytes[8])<<48 +
		int64(bytes[9])<<56

	// The hashes and keys are just copy operations
	copy(b.Difficulty[:], bytes[10:10+hash.Len])
	copy(b.Previous[:], bytes[42:hash.Len+42])
	copy(b.PublicKey[:], bytes[74:ed25519.PublicKeySize+74])

	return
}

func (b *Block) Length() (l int) { return int(unsafe.Sizeof(b)) }

func (b *Block) ID() string { return Name }

// PoWHash returns the Proof of Work hash for a given block. We would not use
// this function in a miner because only the timestamp and previous needs to be
// changed between attempts, and would be faster to directly change the bytes.
func (b *Block) PoWHash() (h hash.Hash, err error) {

	var bytes []byte
	if bytes, err = b.Marshal(); log.E.Chk(err) {
		return
	}

	copy(h[:], proof.DivHash4(bytes[:]))
	return
}

// IndexHash is the hash that is used to store and search for blocks, as it is
// far faster to calculate than the PoWHash.
func (b *Block) IndexHash() (h hash.Hash, err error) {

	var bytes []byte
	if bytes, err = b.Marshal(); log.E.Chk(err) {

		return
	}

	copy(h[:], proof.Blake3(bytes))
	return
}

// GetBlock does nothing interesting in this implementation, but in derived,
// extended block formats with extra fields, this would return the embedded
// Block that is extended from.
func (b *Block) GetBlock() *Block { return b }
