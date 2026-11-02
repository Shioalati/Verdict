package hash

import "verdict/internal/domain/hash"

var factories = make(map[string]func() hash.Hasher)

func RegisterFactory(name string, factory func() hash.Hasher) {
	factories[name] = factory
}
