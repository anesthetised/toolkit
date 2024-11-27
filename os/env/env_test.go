package env

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	var (
		testKey = "test"
		err     error
	)

	t.Run("string", func(t *testing.T) {
		const testValue = "test"

		err = os.Setenv(testKey, testValue)
		assert.NoError(t, err)

		val := Get[string](testKey, "")
		assert.Equal(t, testValue, val)
	})

	t.Run("bool", func(t *testing.T) {
		const testValue = true

		err = os.Setenv(testKey, "true")
		assert.NoError(t, err)

		val := Get[bool](testKey, false)
		assert.Equal(t, testValue, val)
	})

	t.Run("unsigned", func(t *testing.T) {
		t.Run("uint", func(t *testing.T) {
			const testValue = uint(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uint](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("uint8", func(t *testing.T) {
			const testValue = uint8(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uint8](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("uint16", func(t *testing.T) {
			const testValue = uint16(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uint16](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("uint32", func(t *testing.T) {
			const testValue = uint32(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uint32](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("uint64", func(t *testing.T) {
			const testValue = uint64(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uint64](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("uintptr", func(t *testing.T) {
			const testValue = uintptr(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[uintptr](testKey, 0)
			assert.Equal(t, testValue, val)
		})
	})

	t.Run("signed", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			const testValue = int(1)

			err = os.Setenv(testKey, strconv.Itoa(testValue))
			assert.NoError(t, err)

			val := Get[int](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("int8", func(t *testing.T) {
			const testValue = int8(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[int8](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("int16", func(t *testing.T) {
			const testValue = int16(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[int16](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("int32", func(t *testing.T) {
			const testValue = int32(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[int32](testKey, 0)
			assert.Equal(t, testValue, val)
		})

		t.Run("int64", func(t *testing.T) {
			const testValue = int64(1)

			err = os.Setenv(testKey, strconv.Itoa(int(testValue)))
			assert.NoError(t, err)

			val := Get[int64](testKey, 0)
			assert.Equal(t, testValue, val)
		})
	})

	t.Run("float32", func(t *testing.T) {
		const testValue = float32(1)
		err = os.Setenv(testKey, strconv.FormatFloat(float64(testValue), 'f', -1, 32))
		assert.NoError(t, err)

		val := Get[float32](testKey, 0)
		assert.Equal(t, testValue, val)
	})

	t.Run("float64", func(t *testing.T) {
		const testValue = float32(1)
		err = os.Setenv(testKey, strconv.FormatFloat(float64(testValue), 'f', -1, 64))
		assert.NoError(t, err)

		val := Get[float64](testKey, 0)
		assert.Equal(t, testValue, val)
	})

	t.Run("time.Duration", func(t *testing.T) {
		const testValue time.Duration = 30 * time.Second

		err = os.Setenv(testKey, testValue.String())
		assert.NoError(t, err)

		val := Get[time.Duration](testKey, 0)
		assert.Equal(t, testValue, val)
	})
}
