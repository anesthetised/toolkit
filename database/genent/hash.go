package genent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HashField is a wrapper around field.String
// returning an immutable non-empty string with length of 64.
func HashField(name string, unique bool) ent.Field {
	const hashLen = 64

	sb := field.String(name).
		NotEmpty().
		MinLen(hashLen).
		MaxLen(hashLen).
		Immutable()

	if unique {
		sb = sb.Unique()
	}

	return sb
}
