package initialize

import (
	"strings"
)

func Handle(s File) error {
	s.Save("verdict.json", strings.NewReader(""))

	return nil
}
