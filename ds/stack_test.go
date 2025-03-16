package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New[int]()

	t.Run("initial size is zero", func(t *testing.T) {
		assert.Equal(t, 0, s.Len())
	})

	t.Run("push onto the stack", func(t *testing.T) {
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Len())
	})

	t.Run("pop item from the stack", func(t *testing.T) {
		val, ok := s.Pop()
		assert.Equal(t, true, ok)
		assert.Equal(t, 2, val)

		assert.Equal(t, 1, s.Len())

		val, ok = s.Pop()
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, val)

		assert.Equal(t, 0, s.Len())

		val, ok = s.Pop()
		assert.Equal(t, false, ok)
		assert.Equal(t, 0, val)
	})
}
