// Package di facilitates dependency injection with context.Context as a container.
// It relies on reflection to derive a context key from a type.
package di

import (
	"context"
	"reflect"
	"strings"
)

type contextKey string

const defaultPrefix = "di"

// deriveKeyFromType derives a context.Context key from the value's type, concatenating it with the given prefix.
func deriveKeyFromType(prefix string, typ reflect.Type) contextKey {
	s := []string{defaultPrefix, prefix, typ.PkgPath(), typ.Name(), typ.String()}
	return contextKey(strings.TrimLeft(strings.Join(s, "."), "."))
}

// PutNamed returns a new context with a provided value.
// Name is an optional parameter.
func PutNamed[T any](ctx context.Context, name string, val T) context.Context {
	key := deriveKeyFromType(name, reflect.TypeOf(val))
	return context.WithValue(ctx, key, val)
}

// Put is an alias to PutNamed.
func Put[T any](ctx context.Context, val T) context.Context {
	return PutNamed[T](ctx, "", val)
}

// GetNamed retrieves a value from a context.
// Name is an optional parameter.
// It returns a boolean value to indicate a desirable value's presence.
func GetNamed[T any](ctx context.Context, name string) (T, bool) {
	var zero T
	key := deriveKeyFromType(name, reflect.TypeOf(zero))
	val, ok := ctx.Value(key).(T)
	return val, ok
}

// Get is an alias to GetNamed.
func Get[T any](ctx context.Context) (T, bool) {
	return GetNamed[T](ctx, "")
}
