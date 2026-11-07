package root

import "io"

type File interface {
	Save(path string, r io.Reader) (err error)
	Read(path string) (r io.Reader, err error)

	IsExist(path string) (bool, error)
	IsDir(path string) (bool, error)
}
