package hashing

import "errors"

type Hasher interface {
	Hash(data []byte) ([]byte, error)
	Verify(hash []byte, data []byte) error
}

var ErrHashMismatch = errors.New("hash does not match the data")
