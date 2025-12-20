package hashing

type SHA256Hasher struct{}

var _ Hasher = (*SHA256Hasher)(nil)

func NewSHA256Hasher() *SHA256Hasher {
	return &SHA256Hasher{}
}
