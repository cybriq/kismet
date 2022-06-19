package chain

import (
	"unsafe"

	"github.com/cybriq/kismet/pkg/marshal"
	"github.com/dgraph-io/badger/v3"
)

type Chain struct {
	*badger.DB
}

// Link stores the reference between a given block number and its parent
type Link struct {
	Self, Prev uint64
}

const Name = "ChainLink"

var _ marshal.Marshaler = &Link{}

var LinkSize = func() int {
	return int(unsafe.Sizeof(Link{}))
}()

func (link Link) Marshal() (bytes []byte, err error) {

	return
}

func (link *Link) Unmarshal(bytes []byte) (err error) {

	return
}

func (link Link) Length() (l int) {

	return LinkSize
}

func (link Link) ID() string {
	return Name
}
