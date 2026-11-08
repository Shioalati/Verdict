package root

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Shioalati/Verdict/internal/domain/hash"
)

func Handle(l File, h hash.Hasher) (b []byte, err error) {
	isExist, err := l.IsExist(Path)
	if isExist == false || err != nil {
		return nil, err
	}

	isdir, err := l.IsDir(Path)
	if err != nil {
		return nil, err
	}

	if isdir {
		return dir(l, h)
	}

	return notdir(l, h)
}

func dir(l File, h hash.Hasher) (b []byte, err error) {
	if List {
		return list(l, h)
	}

	var buf bytes.Buffer
	err = l.Walk(Path, func(path, relPath string, isDir bool) error {
		if isDir {
			return nil
		}

		sum, err := hashFile(l, h, path)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(&buf, "%s\x00%x\n", relPath, sum)
		return err
	})
	if err != nil {
		return nil, err
	}

	sum, err := h.Hash(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	return []byte(hex.EncodeToString(sum)), nil
}

func list(l File, h hash.Hasher) (b []byte, err error) {
	var lines []string
	err = l.Walk(Path, func(path, relPath string, isDir bool) error {
		if isDir {
			return nil
		}

		sum, err := hashFile(l, h, path)
		if err != nil {
			return err
		}

		lines = append(lines, fmt.Sprintf("%x  %s", sum, relPath))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return []byte(strings.Join(lines, "\n")), nil
}

func hashFile(l File, h hash.Hasher, path string) ([]byte, error) {
	f, err := l.Read(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return h.Hash(f)
}

func notdir(l File, h hash.Hasher) (b []byte, err error) {
	f, err := l.Read(Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result, err := h.Hash(f)
	if err != nil {
		return nil, err
	}

	return []byte(hex.EncodeToString(result)), nil
}
