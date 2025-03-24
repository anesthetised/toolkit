package toolkit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		const expected = "test"
		assert.Equal(t, expected, Must(func() (string, error) { return expected, nil }()))
		assert.Panics(t, func() { Must(func() (string, error) { return "", errors.New(expected) }()) })
	})

	t.Run("bool", func(t *testing.T) {
		const expected = false
		assert.Equal(t, expected, Must(func() (bool, bool) { return expected, true }()))
		assert.Panics(t, func() { Must(func() (bool, bool) { return expected, expected }()) })
	})

	t.Run("unexpected type", func(t *testing.T) {
		assert.Panics(t, func() { Must(func() (bool, struct{}) { return true, struct{}{} }()) })
	})
}
