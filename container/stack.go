package container

type StackErr string

func (s StackErr) Error() string {
	return string(s)
}

const (
	ErrEmpty = StackErr("Stack is empty")
)

type Stack[T any] struct {
	store []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		store: make([]T, 0),
	}
}

func (s *Stack[T]) Empty() bool {
	return len(s.store) == 0
}

func (s *Stack[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Stack[T]) Top() (T, error) {
	if s.Empty() {
		var zero T
		return zero, ErrEmpty
	}

	topIndex := len(s.store) - 1
	return s.store[topIndex], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var zero T
		return zero, ErrEmpty
	}

	topIndex := len(s.store) - 1
	res := s.store[topIndex]
	s.store = s.store[:topIndex]
	return res, nil
}
