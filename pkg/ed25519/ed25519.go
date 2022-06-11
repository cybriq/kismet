// Package ed25519 is a wrapper around golang. org/x/crypto/ed25519 that uses
// arrays instead of indeterminate sized slices to simplify block structure
package ed25519

import (
	"fmt"
	ed "golang.org/x/crypto/ed25519"
	"io"
)

const PublicKeySize = ed.PublicKeySize

type PublicKey [PublicKeySize]byte

func (pk PublicKey) Equal(x PublicKey) bool {
	return ed.PublicKey(pk[:]).Equal(ed.PublicKey(x[:]))
}

type Seed [ed.SeedSize]byte

type PrivateKey [ed.PrivateKeySize]byte

func (sk PrivateKey) Public() (pk PublicKey) {
	copy(pk[:], ed.PrivateKey(sk[:]).Public().(ed.PrivateKey))
	return
}

func (sk PrivateKey) Equal(x PrivateKey) bool {
	return sk == x
}

func (sk PrivateKey) Seed() (seed Seed) {
	copy(seed[:], sk[:32])
	return
}

func GenerateKey(rr io.Reader) (pk PublicKey, sk PrivateKey, err error) {

	p, s, er := ed.GenerateKey(rr)

	if er != nil {
		err = er
		return
	}

	copy(pk[:], p)
	copy(sk[:], s)
	return
}

func NewKeyFromSeed(s Seed) (pk PrivateKey) {

	p := ed.NewKeyFromSeed(s[:])
	copy(pk[:], p)
	return
}

type Signature [ed.SignatureSize]byte

func (sk PrivateKey) Sign(message []byte) (sig Signature) {

	s := ed.Sign(sk[:], message)
	copy(sig[:], s)
	return
}

func (pk PublicKey) Verify(message []byte, sig Signature) bool {
	return ed.Verify(pk[:], message, sig[:])
}

func ToSignature(s []byte) (sig Signature, err error) {
	if len(s) != ed.SignatureSize {
		err = fmt.Errorf("incorrect signature size, got %d expected %d", len(s), ed.SignatureSize)
		return
	}

	copy(sig[:], s)
	return
}
