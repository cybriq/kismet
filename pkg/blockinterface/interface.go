package blockinterface

import (
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
)

// Type is a generic interface for general purpose handling of blocks between
// their wire and in memory formats.
//
// Because our functions are assumed to mutate the implemented version of a
// block, we can create a Deserialize function because the type is the
// interface. An implementation of this interface must require an allocated data
// structure to copy into, as a pointer type with nil cannot mutate its contents
// and thus fails to implement the interface strictly.
//
// This interface is being created because the base block.Block type has
// multiple types and implementing them is intended to require creating a new
// interface so that using them generically as raw bytes does not require
// concrete handling, and unwrapping bytes into a concrete type does not require
// type assertion.
//
// The generic Hash functions are here for generic use. Performance is
// sacrificed for simplicity in these functions, miners would work with the
// concrete types directly for such things as changing timestamps and previous
// blocks.
//
// Embedding can be used for convenience to access the fields of the type
// directly. It is intended that the core block.Block type is embedded into
// extended versions, and thus the concrete type will be accessible by accessing
// the .Block field of derivatives. Thus, this getter is in the interface. For
// extended fields, this interface will assume the consuming code knows the type
// to assert to.
//
// This also implicitly dictates that any implementation of this interface must
// be derived from the block.Block type. The type contains purely the generic
// elements of a block, without any payload. Most likely extensions will contain
// only one extra hash, being a reference to another data repository of some
// kind, but it could be more than this, so it is left open.
//
// Actual handling of unknown block types will be handled in a separate signed
// packet library that deals only with signed raw bytes, which will also assume
// the first two bytes of the payload are a type identifier and provide an
// interface to add to block code that checks a table of known types, which will
// be found in the block package anyway.
type Type interface {
	SerialLen() int
	Serialize() (bytes []byte, err error)
	Deserialize(bytes []byte) (err error)
	PoWHash() (h hash.Hash, err error)
	IndexHash() (h hash.Hash, err error)
	GetBlock() *block.Block
}
