package ds

type node[T any] struct {
	val  T
	prev *node[T]
}

// Stack implements a generic stack implementation
// which is shamelessly copied from https://github.com/golang-collections/collections.
type Stack[T any] struct {
	top *node[T]
	len int
}

// New returns a new Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Len returns a number of items on the stack.
func (s *Stack[T]) Len() int {
	return s.len
}

// Push places an item onto the top of the stack.
func (s *Stack[T]) Push(val T) {
	n := &node[T]{val: val, prev: s.top}
	s.top = n
	s.len++
}

// Pop retrieves an item from the top of the stack and returns it alongside a boolean value
// which indicates whether an item is present.
func (s *Stack[T]) Pop() (T, bool) {
	if s.len == 0 {
		var zero T
		return zero, false
	}

	n := s.top
	s.top = n.prev
	s.len--

	return n.val, true
}
