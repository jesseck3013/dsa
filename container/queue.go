package container

type QueueErr string

func (s QueueErr) Error() string {
	return string(s)
}

const (
	ErrQueueEmpty = QueueErr("Queue is empty")
)

type Queue[T any] struct {
	store []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		store: make([]T, 0),
	}
}

func (q *Queue[T]) Empty() bool {
	return len(q.store) == 0
}

func (q *Queue[T]) Enqueue(value T) {
	q.store = append(q.store, value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Empty() {
		var zero T
		return zero, ErrQueueEmpty
	}

	res := q.store[0]
	q.store = q.store[1:]
	return res, nil
}
