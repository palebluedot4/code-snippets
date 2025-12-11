package hashing

import "errors"

var ErrHashMismatch = errors.New("hash does not match the data")

type Hasher interface {
	Hash(data []byte) ([]byte, error)
	Verify(hash []byte, data []byte) error
}
