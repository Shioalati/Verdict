package hash

import (
	"crypto/sha256"
	"io"
	"verdict/internal/domain/hash"
)

func init() {
	RegisterFactory("sha256", func() hash.Hasher {
		return &SHA256{}
	})
}

type SHA256 struct{}

func (*SHA256) Hash(data io.Reader) ([]byte, error) {
	h := sha256.New()
	_, err := io.Copy(h, data)
	return h.Sum(nil), err
}
