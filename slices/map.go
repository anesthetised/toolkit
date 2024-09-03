package slices

// Map provides a convenient way to map slice into a different type.
// It always returns a non-nil slice.
func Map[T, U any](s []T, fn func(i int, v T) U) []U {
	ms := []U{}

	for i, v := range s {
		ms = append(ms, fn(i, v))
	}

	return ms
}
