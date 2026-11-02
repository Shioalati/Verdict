package hash

import "io"

type Hasher interface {
	Hash(data io.Reader) ([]byte, error)
}
