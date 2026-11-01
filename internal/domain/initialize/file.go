package initialize

import "io"

type File interface {
	Save(path string, r io.Reader) (err error)
}
