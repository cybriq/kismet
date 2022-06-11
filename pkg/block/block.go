package block

import (
	"fmt"
	"github.com/cybriq/kismet/pkg/ed25519"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/known"
	"github.com/cybriq/kismet/pkg/proof"
	"unsafe"
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

// GetBlock does nothing interesting in this implementation, but in derived,
// extended block formats with extra fields, this would return the embedded
// Block that is extended from.
func (b *Block) GetBlock() *Block { return b }

// SerialLen returns the length in bytes of the Marshal ed version
func (b Block) SerialLen() int { return int(unsafe.Sizeof(b)) }

const WireBlockLen = 2 + 8 + hash.Len*2 + ed25519.PublicKeySize

// WireBlock is defined here as an array as this simplifies data validation
type WireBlock [WireBlockLen]byte

// Marshal returns the raw bytes for the wire format of the block
func (b *Block) Marshal() (serial WireBlock, err error) {

	if b == nil {
		// It is programmer error if a nil pointer is passed
		err = fmt.Errorf("cannot marshal nil Block")
		return
	}

	// There is functions to do these but that would be slower than doing this
	// directly with the integers
	serial[0] = byte(b.Type)
	serial[1] = byte(b.Type >> 8)
	serial[2] = byte(b.Time)
	serial[3] = byte(b.Time >> 8)
	serial[4] = byte(b.Time >> 16)
	serial[5] = byte(b.Time >> 24)
	serial[6] = byte(b.Time >> 32)
	serial[7] = byte(b.Time >> 40)
	serial[8] = byte(b.Time >> 48)
	serial[9] = byte(b.Time >> 56)

	// The rest are simple copy operations
	copy(serial[10:hash.Len+10], b.Difficulty[:])
	copy(serial[42:hash.Len+42], b.Previous[:])
	copy(serial[74:ed25519.PublicKeySize+74], b.PublicKey[:])
	return
}

// Serialize renders our block into wire/storage format
func (b *Block) Serialize() (bytes []byte, err error) {

	var wb WireBlock

	wb, err = b.Marshal()
	if err != nil {
		return
	}

	bytes = wb[:]
	return
}

func ToWireBlock(b []byte) (wb WireBlock, err error) {
	if len(b) != WireBlockLen {
		err = fmt.Errorf(
			"data length incorrect, got %d expected %d",
			len(b), WireBlockLen,
		)
		return
	}

	copy(wb[:], b)
	return
}

// Unmarshal the wire format bytes into a Block structure
func (serial WireBlock) Unmarshal() (b Block) {

	// again, just doing this directly is the fastest.
	b.Type = known.Type(serial[0]) + known.Type(serial[1])<<8
	b.Time = int64(serial[2]) +
		int64(serial[3])<<8 +
		int64(serial[4])<<16 +
		int64(serial[5])<<24 +
		int64(serial[6])<<32 +
		int64(serial[7])<<40 +
		int64(serial[8])<<48 +
		int64(serial[9])<<56

	// The hashes and keys are just copy operations
	copy(b.Difficulty[:], serial[10:10+hash.Len])
	copy(b.Previous[:], serial[42:hash.Len+42])
	copy(b.PublicKey[:], serial[74:ed25519.PublicKeySize+74])
	return
}

// Deserialize unpacks a serialized form of the block into the value referred to by the pointer
func (b *Block) Deserialize(bytes []byte) (err error) {

	if b == nil {

		err = fmt.Errorf("cannot deserialize to a nil Block")
		return
	}

	var wb WireBlock
	if len(bytes) != len(wb) {
		err = fmt.Errorf(
			"cannot deserialize %d bytes as a block is %d bytes long",
			len(bytes), len(wb),
		)
		return
	}
	copy(wb[:], bytes)

	// the assignment here is a copy operation that overwrites the existing Block
	*b = wb.Unmarshal()
	return
}

// PoWHash returns the Proof of Work hash for a given block. We would not use
// this function in a miner because only the timestamp and previous needs to be
// changed between attempts, and would be faster to directly change the bytes.
func (b *Block) PoWHash() (h hash.Hash, err error) {

	var bytes WireBlock
	bytes, err = b.Marshal()
	if err != nil {
		return
	}

	copy(h[:], proof.DivHash4(bytes[:]))
	return
}

// IndexHash is the hash that is used to store and search for blocks, as it is
// far faster to calculate than the PoWHash.
func (b *Block) IndexHash() (h hash.Hash, err error) {

	var bytes WireBlock
	bytes, err = b.Marshal()
	if err != nil {

		return
	}

	copy(h[:], proof.Blake3(bytes[:]))
	return
}
