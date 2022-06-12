// Package ed25519 is a wrapper around golang. org/x/crypto/ed25519 that uses
// arrays instead of indeterminate sized slices to simplify block structure
package ed25519

import (
	"fmt"
	ed "golang.org/x/crypto/ed25519"
	"io"
)

// PublicKeySize is the size in bytes of a PublicKey
const PublicKeySize = ed.PublicKeySize

// PublicKey is an ed25519 PublicKey
type PublicKey [PublicKeySize]byte

// Equal returns true if another PublicKey is the same as the current one.
func (pk PublicKey) Equal(x PublicKey) bool {
	return ed.PublicKey(pk[:]).Equal(ed.PublicKey(x[:]))
}

// SeedSize is the size of an ed25519 Seed
const SeedSize = ed.SeedSize

// Seed is a seed for a ed25519 key pair
type Seed [SeedSize]byte

// PrivateKey is an ed25519 PrivateKey
type PrivateKey [ed.PrivateKeySize]byte

// Public returns the public key corresponding to a PrivateKey
func (sk PrivateKey) Public() (pk PublicKey) {
	copy(pk[:], ed.PrivateKey(sk[:]).Public().(ed.PrivateKey))
	return
}

// Equal returns true if another PrivateKey is the same. Because we use arrays
// in this library we don't need to call the supporting ed25519 library for this
// simple comparison.
func (sk PrivateKey) Equal(x PrivateKey) bool {
	return sk == x
}

// Seed returns the seed value that can generate the PrivateKey again using NewKeyFromSeed
func (sk PrivateKey) Seed() (seed Seed) {
	copy(seed[:], sk[:32])
	return
}

// GenerateKey creates a new key with an optional random bytes source and
// returns a PrivateKey and matching PublicKey
func GenerateKey(rr io.Reader) (pk PublicKey, sk PrivateKey, err error) {

	p, s, er := ed.GenerateKey(rr)

	if log.E.Chk(er) {
		err = er
		return
	}

	copy(pk[:], p)
	copy(sk[:], s)
	return
}

// NewKeyFromSeed creates a new key based on a Seed
func NewKeyFromSeed(s Seed) (pk PrivateKey) {

	p := ed.NewKeyFromSeed(s[:])
	copy(pk[:], p)
	return
}

// Signature is an ed25519 Signature
type Signature [ed.SignatureSize]byte

// Sign signs a message and returns a Signature
func (sk PrivateKey) Sign(message []byte) (sig Signature) {

	s := ed.Sign(sk[:], message)
	copy(sig[:], s)
	return
}

func (pk PublicKey) Verify(message []byte, sig Signature) bool {
	return ed.Verify(pk[:], message, sig[:])
}

// ToSignature converts bytes presumed to be a signature into a Signature
func ToSignature(s []byte) (sig Signature, err error) {

	if len(s) != ed.SignatureSize {
		err = fmt.Errorf("incorrect signature size, got %d expected %d", len(s), ed.SignatureSize)
		log.E.Ln(err)
		return
	}

	copy(sig[:], s)
	return
}
