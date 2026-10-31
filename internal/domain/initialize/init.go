package initialize

import (
	"strings"
	"verdict/internal/infrastructure/file"
)

func Handle(s *file.LocalStorage) error {
	s.Save("verdict.json", strings.NewReader(""))

	return nil
}
