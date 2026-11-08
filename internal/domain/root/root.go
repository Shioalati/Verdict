package root

import "github.com/Shioalati/Verdict/internal/domain/hash"

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
		dir()
	} else {
		b, err = notdir(l, h)
		return
	}

	return nil, err
}

func dir() {

}

func notdir(l File, h hash.Hasher) (b []byte, err error) {
	file, err := l.Read(Path)
	if err != nil {
		return nil, err
	}

	result, err := h.Hash(file)
	if err != nil {
		return nil, err
	}
	return result, nil
}
