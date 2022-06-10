// Package known is a compile time registry for keeping track of block type
// codes known to the current implementation.
package known

const (
	// ValidatorBlock is for a slot in the block production schedule
	ValidatorBlock Type = iota

	// Unknown is an invalid block type and marks the type number after the last
	// valid type. When more types are added, they should have their variant in a
	// new named type with Type implemented.
	Unknown
)

// Type is the type of a block. Using 16 bit values we can have up to 65536
// which should be more than enough for one platform
type Type uint16
