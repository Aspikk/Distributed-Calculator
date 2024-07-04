package stack

type node[T any] struct {
	value    T
	previous *node[T]
}

type Stack[T any] struct {
	top *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		top: nil,
	}
}

func (s *Stack[T]) Push(value T) {
	new := &node[T]{
		value:    value,
		previous: s.top,
	}

	s.top = new
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var nul T
		return nul, false
	}

	value := s.top.value
	s.top = s.top.previous

	return value, true
}

func (s *Stack[T]) GetTop() T {
	return s.top.value
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}
