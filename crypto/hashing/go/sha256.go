package hashing

import (
	"crypto/sha256"
	"crypto/subtle"
)

type SHA256Hasher struct{}

var _ Hasher = (*SHA256Hasher)(nil)

func NewSHA256Hasher() *SHA256Hasher {
	return &SHA256Hasher{}
}

func (h *SHA256Hasher) Hash(data []byte) ([]byte, error) {
	sum := sha256.Sum256(data)
	return sum[:], nil
}

func (h *SHA256Hasher) Verify(hash []byte, data []byte) error {
	sum := sha256.Sum256(data)
	if subtle.ConstantTimeCompare(hash, sum[:]) != 1 {
		return ErrHashMismatch
	}
	return nil
}
