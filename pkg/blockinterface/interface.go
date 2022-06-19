package block

import (
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/marshal"
)

// Interface is a generic interface for general purpose handling of blocks between
// their wire and in memory formats.
type Interface interface {
	marshal.Marshaler
	PoWHash() (h hash.Hash, err error)
	IndexHash() (h hash.Hash, err error)
	GetBlock() *block.Block
}
