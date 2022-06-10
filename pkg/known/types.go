// Package known is a compile time registry for keeping track of block type
// codes known to the current implementation.
package known

import (
	"github.com/cybriq/kismet/pkg/block"
)

const (
	// ValidatorBlock is for a slot in the block production schedule
	ValidatorBlock block.BlockType = iota

	// NullBlock is an invalid block type and marks the type number after the last
	// valid type. When more types are added, they should have their variant in a
	// new named type with blockinterface.Type implemented.
	NullBlock
)
