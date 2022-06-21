package nodes

import (
	"fmt"
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

func NewNode() *Node { return &Node{} }

var FixedPart = func() int {
	var l Node
	return int(unsafe.Sizeof(l.Number) + unsafe.Sizeof(l.Hash))
}()

var MinBigInt = func() int {
	return len(big.NewInt(0).Bytes())
}()

const Name = "nodes.Node"

var _ marshal.Marshaler = &Node{}

func (n Node) Marshal() (bytes []byte, err error) {

	return
}

func (n *Node) Unmarshal(bytes []byte) (err error) {

	if len(bytes) < FixedPart+MinBigInt {
		err = fmt.Errorf(
			"data length less than minimum, got %d expected minimum %d",
			len(bytes), FixedPart+MinBigInt,
		)
		log.E.Ln(err)
		return
	}

	return
}

func (n Node) Length() (l int) {

	return FixedPart + n.Weight.BitLen()
}

func (n Node) ID() string { return Name }
