package hashing

type BLAKE2bHasher struct{}

var _ Hasher = (*BLAKE2bHasher)(nil)

func NewBLAKE2bHasher() *BLAKE2bHasher {
	return &BLAKE2bHasher{}
}
