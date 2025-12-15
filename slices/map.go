package slices

// Map applies fn to each element of s and returns a new slice of the results.
// The order of elements is preserved. The returned slice has length len(s) and
// capacity cap(s). If s is nil, Map returns an empty (non-nil) slice.
func Map[T, U any](s []T, fn func(i int, v T) U) []U {
	ms := make([]U, len(s), cap(s))

	for i, v := range s {
		ms[i] = fn(i, v)
	}

	return ms
}
