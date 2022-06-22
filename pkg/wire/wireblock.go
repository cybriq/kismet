// Package wireblock provides functions for directly interacting with already
// marshalled blocks such as updating timestamp, changing previous block hash
// and public key, for use by miners most specifically.
package wire

import (
	"fmt"
	"time"

	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/ed25519"
	"github.com/cybriq/kismet/pkg/hash"
)

// Block is an array so that size and bounds checks are automatically
// handled, to hash it must be converted using FromWireBlock
type Block [block.Length]byte

// ToWireBlock converts the bytes from block.Marshal into the Block array
func ToWireBlock(by []byte) (b *Block, err error) {

	if len(by) != block.Length {
		err = fmt.Errorf("cannot convert block: incorrect byte length: "+
			"got %d expected %d", len(by), block.Length)
		log.E.Chk(err)
		return
	}

	b = &Block{}
	copy(b[:], by)

	return
}

// UpdateTime sets the block's Time field to the current time
func (b *Block) UpdateTime() {
	t := time.Now().UnixNano()
	b[2] = byte(t)
	b[3] = byte(t >> 8)
	b[4] = byte(t >> 16)
	b[5] = byte(t >> 24)
	b[6] = byte(t >> 32)
	b[7] = byte(t >> 40)
	b[8] = byte(t >> 48)
	b[9] = byte(t >> 56)
}

// SetDifficulty sets the block.Block.Difficulty for the block
func (b *Block) SetDifficulty(d hash.Hash) { copy(b[10:42], d[:]) }

// SetPrevious changes the block.Block.Previous field
func (b *Block) SetPrevious(p hash.Hash) { copy(b[42:74], p[:]) }

// SetPublicKey changes the block.Block.PublicKey field
func (b *Block) SetPublicKey(k ed25519.PublicKey) { copy(b[74:106], k[:]) }

// FromWireBlock slices the array back to []byte for use with hash functions
// and network sends.
func (b *Block) FromWireBlock() []byte { return b[:] }
