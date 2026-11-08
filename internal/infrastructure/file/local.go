package file

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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

func (l LocalStorage) Read(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func (l LocalStorage) IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func (l LocalStorage) IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func (l LocalStorage) Walk(root string, fn func(path, relPath string, isDir bool) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Type()&os.ModeSymlink != 0 {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		return fn(path, relPath, d.IsDir())
	})
}
