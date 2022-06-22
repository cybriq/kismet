package blockinterface

import (
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/kismet/pkg/marshal"
)

// Blocker is a generic interface for general purpose handling of blocks between
// their wire and in memory formats.
type Blocker interface {
	marshal.Marshaler
	PoWHash() (h []byte, err error)
	IndexHash() (h hash.Hash, err error)
	GetBlock() *block.Block
}
