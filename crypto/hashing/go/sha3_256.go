package hashing

type SHA3256Hasher struct{}

var _ Hasher = (*SHA3256Hasher)(nil)

func NewSHA3256Hasher() *SHA3256Hasher {
	return &SHA3256Hasher{}
}
