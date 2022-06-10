// Package packet is a generic wrapper for containing messages that are most
// likely blocks for kismet's proof of work chains, used to enable generic
// processing of messages to recognise if they are known or not.
//
// This will be used by the p2p system as a first step for decoding packets so
// that potentially unknown but signed messages can be ignored while processing
// known messages, essential for extensibility without forcing immediate
// upgrade.
package packet

import (
	"encoding/binary"
	"fmt"
)

// UnmarshalGeneric unpacks the contents of a wire format with an int64 length
// prefix and uint16 type as the first field, checks the length and returns the
// type number. With this, a generic packet format is created with a 16 bit
// identifier.
//
// The excess field would usually contain a signature, thus enabling the generic
// processing of signed packets where unknown types can be expected to have a
// signature that must be valid.
func UnmarshalGeneric(pkt []byte) (typ uint16, bytes []byte, excess []byte, err error) {

	if len(pkt) < 10 {
		err = fmt.Errorf("packet does not contain length and type")
		return
	}

	pktSize, _ := binary.Varint(pkt[:8])
	if int64(len(pkt)-8) < pktSize {
		err = fmt.Errorf("packet length should be %d but %d bytes given", pktSize, len(pkt)-8)
		return
	}
	// This enables generic packaging of signed packets
	if int64(len(pkt)-8) > pktSize {
		excess = pkt[:pktSize]
	}
	typ = uint16(pkt[8]) + uint16(pkt[9])>>8
	bytes = pkt[8:]
	return
}

// MarshalGeneric simply appends the length to raw bytes. It is the
// responsibility of the implementation using this to encode the type as the
// first two bytes.
//
// This is not intended to be used with large data packets as it copies the
// packet into a second slice.
//
// For the case of signed packets, the sender would then sign this output and append.
func MarshalGeneric(bytes []byte) (out []byte) {

	out = make([]byte, len(bytes)+8)
	binary.PutVarint(out[:8], int64(len(bytes)))
	copy(out[8:], bytes)
	return
}
