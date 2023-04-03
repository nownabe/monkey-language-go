package object

import "hash/fnv"

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	hk := HashKey{Type: b.Type()}

	if b.Value {
		hk.Value = 1
	} else {
		hk.Value = 0
	}

	return hk
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
