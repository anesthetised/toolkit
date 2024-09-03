package ds

// Pair implements a coupled key-value data type.
type Pair[K, V any] struct {
	key   K
	value V
}

// NewPair returns a new Pair.
func NewPair[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{
		key:   key,
		value: value,
	}
}

// Values returns a key and an associated value.
func (p Pair[K, V]) Values() (K, V) {
	return p.key, p.value
}
