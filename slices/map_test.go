package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("slice is not nil", func(t *testing.T) {
		assert.NotNil(t, Map(nil, func(_ int, v int) int { return v }))
	})

	t.Run("len and cap are preserved", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := Map(a, func(_ int, v int) int { return v + 2 })
		assert.Equal(t, len(a), len(b))
		assert.Equal(t, cap(a), cap(b))
		assert.NotEqual(t, a, b)
	})
}
