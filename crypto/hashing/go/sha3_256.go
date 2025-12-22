package hashing

import (
	"crypto/subtle"

	"golang.org/x/crypto/sha3"
)

type SHA3256Hasher struct{}

var _ Hasher = (*SHA3256Hasher)(nil)

func NewSHA3256Hasher() *SHA3256Hasher {
	return &SHA3256Hasher{}
}

func (h *SHA3256Hasher) Hash(data []byte) ([]byte, error) {
	sum := sha3.Sum256(data)
	return sum[:], nil
}

func (h *SHA3256Hasher) Verify(hash []byte, data []byte) error {
	sum := sha3.Sum256(data)
	if subtle.ConstantTimeCompare(hash, sum[:]) != 1 {
		return ErrHashMismatch
	}
	return nil
}
