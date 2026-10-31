package file

import (
	"fmt"
	"io"
	"os"
)

type LocalStorage struct{}

func NewLocal() *LocalStorage {
	return &LocalStorage{}
}

func (l LocalStorage) Save(path string, r io.Reader) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("storage: create file %q: %w", path, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("storage: close file %q: %w", path, cerr)
		}
	}()

	if _, err := io.Copy(f, r); err != nil {
		return fmt.Errorf("storage: write file %q: %w", path, err)
	}

	return nil
}
