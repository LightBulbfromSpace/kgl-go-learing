package generics

type StackOfInts = Stack[int]
type StackOfStrings = Stack[string]

type Stack[T any] struct {
	values []T
}

// Push appends value to the end
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// IsEmpty return true if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Pop removes the last value in stack, returns as the first parameter value of removed item
// and false as the second parameter if stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
