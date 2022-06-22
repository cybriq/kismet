// Package wireblock provides functions for directly interacting with already
// marshalled blocks such as updating timestamp, changing previous block hash
// and public key, for use by miners most specifically.
package wire

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/blockinterface"
	"github.com/cybriq/kismet/pkg/ed25519"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/proof"
	"lukechampine.com/blake3"
)

// Block is an array so that size and bounds checks are automatically
// handled, to hash it must be converted using FromWireBlock
type Block [block.Length]byte

var _ blockinterface.Blocker = Block{}
var Name = reflect.TypeOf(Block{}).Name()

// Unmarshal converts the bytes from block.Marshal into the Block array
func (b Block) Unmarshal(by []byte) (err error) {

	if len(by) != block.Length {
		err = fmt.Errorf("cannot convert block: incorrect byte length: "+
			"got %d expected %d", len(by), block.Length)
		log.E.Chk(err)
		return
	}

	copy(b[:], by)

	return
}

// UpdateTime sets the block's Time field to the current time
func (b Block) UpdateTime() {
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

func (b Block) SetDifficulty(d hash.Hash)        { copy(b[10:42], d[:]) }
func (b Block) SetPrevious(p hash.Hash)          { copy(b[42:74], p[:]) }
func (b Block) SetPublicKey(k ed25519.PublicKey) { copy(b[74:106], k[:]) }
func (b Block) Length() (l int)                  { return block.Length }
func (b Block) ID() string                       { return Name }
func (b Block) Marshal() ([]byte, error)         { return b[:], nil }

func (b Block) PoWHash() (h []byte, err error) {

	h = proof.DivHash4(b[:])
	return
}

func (b Block) IndexHash() (h hash.Hash, err error) {

	h = blake3.Sum256(b[:])
	return

}

func (b Block) GetBlock() (bl *block.Block) {

	// this function only fails with wrong size and size is always correct
	// for array
	_ = bl.Unmarshal(b[:])
	return
}
