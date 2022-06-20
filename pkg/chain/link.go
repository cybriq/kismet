package chain

import (
	"fmt"
	"unsafe"

	"github.com/cybriq/kismet/pkg/bytes"
	"github.com/cybriq/kismet/pkg/marshal"
)

// Link stores the reference between a given block number and its parent
type Link struct {
	Self, Prev uint64
}

func NewLink() *Link { return &Link{} }

var _ marshal.Marshaler = &Link{}

const Name = "chain.Link"

var LinkSize = func() int { return int(unsafe.Sizeof(Link{})) }()

func (link *Link) Marshal() (b []byte, err error) {

	b = make([]byte, LinkSize)

	copy(b[:8], bytes.FromUint64(link.Self))
	copy(b[8:], bytes.FromUint64(link.Prev))

	return
}

func (link *Link) Unmarshal(b []byte) (err error) {

	if len(b) != LinkSize {
		err = fmt.Errorf(
			"data length incorrect, got %d expected %d",
			len(b), LinkSize,
		)
		log.E.Ln(err)
		return
	}

	link.Self = bytes.ToUint64(b[:8])
	link.Prev = bytes.ToUint64(b[8:])
	return
}

func (link Link) Length() (l int) { return LinkSize }
func (link Link) ID() string      { return Name }
