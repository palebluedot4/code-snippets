package hashing

import (
	"crypto/md5"
	"crypto/subtle"
)

type MD5Hasher struct{}

var _ Hasher = (*MD5Hasher)(nil)

func NewMD5Hasher() *MD5Hasher {
	return &MD5Hasher{}
}

func (h *MD5Hasher) Hash(data []byte) ([]byte, error) {
	sum := md5.Sum(data)
	return sum[:], nil
}

func (h *MD5Hasher) Verify(hash []byte, data []byte) error {
	sum := md5.Sum(data)
	if subtle.ConstantTimeCompare(hash, sum[:]) != 1 {
		return ErrHashMismatch
	}
	return nil
}
