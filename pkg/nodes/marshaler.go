package nodes

import (
	"math/big"
	"unsafe"

	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/marshal"
)

type Node struct {
	Number uint64
	Hash   hash.Hash
	Weight big.Int
}

var FixedPart = func() int {
	var l Node
	return int(unsafe.Sizeof(l.Number) + unsafe.Sizeof(l.Hash))
}()

const Name = "BlockNode"

var _ marshal.Marshaler = &Node{}

func (n Node) Marshal() (bytes []byte, err error) {

	return
}

func (n *Node) Unmarshal(bytes []byte) (err error) {

	return
}

func (n Node) Length() (l int) {

	return FixedPart + n.Weight.BitLen()
}

func (n Node) ID() string {
	return Name
}
