// Package env provides a way to automatically convert an environment variable value
// to a variety of types.
package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

func NewEmptyEnvError(name string) error {
	return fmt.Errorf("%s environment variable must not be empty", name)
}

// Get retrieves an environment variable, converting it to the desirable type
// constrained by a limited set of types which are applicable to environment values.
func Get[T any](key string, fallback T) T {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	untyped := any(value)

	switch any(fallback).(type) {
	case string:
		return untyped.(T)
	case bool:
		return any(parseBool(value, any(fallback).(bool))).(T)
	case uint:
		return any(parseUnsigned[uint](value, any(fallback).(uint))).(T)
	case uint8:
		return any(parseUnsigned[uint8](value, any(fallback).(uint8))).(T)
	case uint16:
		return any(parseUnsigned[uint16](value, any(fallback).(uint16))).(T)
	case uint32:
		return any(parseUnsigned[uint32](value, any(fallback).(uint32))).(T)
	case uint64:
		return any(parseUnsigned[uint64](value, any(fallback).(uint64))).(T)
	case uintptr:
		return any(parseUnsigned[uintptr](value, any(fallback).(uintptr))).(T)
	case int:
		return any(parseSigned[int](value, any(fallback).(int))).(T)
	case int8:
		return any(parseSigned[int8](value, any(fallback).(int8))).(T)
	case int16:
		return any(parseSigned[int16](value, any(fallback).(int16))).(T)
	case int32:
		return any(parseSigned[int32](value, any(fallback).(int32))).(T)
	case int64:
		return any(parseSigned[int64](value, any(fallback).(int64))).(T)
	case float32:
		return any(parseFloat[float32](value, any(fallback).(float32))).(T)
	case float64:
		return any(parseFloat[float64](value, any(fallback).(float64))).(T)
	case time.Duration:
		dur, err := time.ParseDuration(value)
		if err != nil {
			return fallback
		}
		return any(dur).(T)
	default:
		return fallback
	}
}

func parseBool(raw string, fallback bool) bool {
	val, err := strconv.ParseBool(raw)
	if err != nil {
		return fallback
	}
	return val
}

func parseSigned[T constraints.Signed](raw string, fallback T) T {
	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return fallback
	}
	return T(val)
}

func parseUnsigned[T constraints.Unsigned](raw string, fallback T) T {
	val, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return fallback
	}
	return T(val)
}

func parseFloat[T constraints.Float](raw string, fallback T) T {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fallback
	}
	return T(val)
}
