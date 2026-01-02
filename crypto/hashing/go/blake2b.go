package hashing

import (
	"crypto/subtle"

	"golang.org/x/crypto/blake2b"
)

type BLAKE2bHasher struct{}

var _ Hasher = (*BLAKE2bHasher)(nil)

func NewBLAKE2bHasher() *BLAKE2bHasher {
	return &BLAKE2bHasher{}
}

func (h *BLAKE2bHasher) Hash(data []byte) ([]byte, error) {
	sum := blake2b.Sum256(data)
	return sum[:], nil
}

func (h *BLAKE2bHasher) Verify(hash []byte, data []byte) error {
	sum := blake2b.Sum256(data)
	if subtle.ConstantTimeCompare(hash, sum[:]) != 1 {
		return ErrHashMismatch
	}
	return nil
}
