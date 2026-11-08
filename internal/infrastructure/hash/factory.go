package hash

import (
	"fmt"

	"github.com/Shioalati/Verdict/internal/domain/hash"
)

var factories = make(map[string]func() hash.Hasher)

func RegisterFactory(name string, factory func() hash.Hasher) {
	factories[name] = factory
}

func New(name string) (hash.Hasher, error) {
	f, ok := factories[name]
	if !ok {
		return nil, fmt.Errorf("unknown hash algorithm: %s", name)
	}

	if name == "" {
		return factories["sha256"](), nil
	}
	return f(), nil
}
