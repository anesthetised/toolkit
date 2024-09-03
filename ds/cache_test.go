package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := NewCache[string, string]()

	const (
		testKey = "test"
		testVal = "test"
	)

	c.Set(testKey, testVal)

	val, ok := c.Get(testKey)
	assert.True(t, ok)
	assert.Equal(t, testVal, val)

	val, ok = c.Get("test123")
	assert.False(t, ok)
	assert.Equal(t, "", val)

	c.Delete(testKey)

	val, ok = c.Get(testKey)
	assert.False(t, ok)
	assert.Equal(t, "", val)
}
