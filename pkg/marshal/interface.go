package marshal

// Marshaler is an interface for any data to be converted to bytes, and back again
type Marshaler interface {
	// Marshal returns generic []byte data. This cannot fail as it should be
	// implemented with a value receiver method.
	Marshal() (bytes []byte, err error)
	// Unmarshal takes bytes and populates itself. The type magic is in the
	// interface, we don't have to specify an output if the type itself is the
	// output. The main error here would be invalid input. In all cases incorrect
	// Length would be at least one of the errors.
	Unmarshal(bytes []byte) (err error)
	// Length returns the byte length of the marshaled form
	Length() (l int)
	// ID returns the name of the type, mainly this is for debugging/logging
	ID() string
}
