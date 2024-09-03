package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	const (
		key   = "key"
		value = "value"
	)

	p := NewPair(key, value)
	k, v := p.Values()

	assert.Equal(t, key, k)
	assert.Equal(t, value, v)
}
