package hashing

import "errors"

var ErrHashMismatch = errors.New("hash does not match the data")

//go:generate go run go.uber.org/mock/mockgen -destination=mock_hasher_test.go -package=hashing -source=hasher.go Hasher
type Hasher interface {
	Hash(data []byte) ([]byte, error)
	Verify(hash []byte, data []byte) error
}
