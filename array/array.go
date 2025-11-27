package array

type StaticArray[T any] struct {
	store  []T
	length uint
}

type ArrayErr string

func (err ArrayErr) Error() string {
	return string(err)
}

const ErrOutOfBound = ArrayErr("index is out of bound")

// Time: O(n)
// Space: O(n)
func NewStaticArray[T any](length uint) *StaticArray[T] {
	store := make([]T, length)
	return &StaticArray[T]{
		store:  store,
		length: length,
	}
}

// Time: O(1)
// Space: O(1)
func (sa *StaticArray[T]) Read(index uint) (T, error) {
	if index >= sa.length {
		var zero T
		return zero, ErrOutOfBound
	}

	return sa.store[index], nil
}

// Time: O(1)
// Space: O(1)
func (sa *StaticArray[T]) Update(index uint, value T) error {
	if index >= sa.length {
		return ErrOutOfBound
	}

	sa.store[index] = value
	return nil
}

// Time: O(1)
// Space: O(1)
// There is no deletion in the data store which means
// the underlying array size does not change.
// This method only shift all items whose index is larger than
// the target index to their left with one index.
func (sa *StaticArray[T]) Delete(index uint) error {
	if index >= sa.length {
		return ErrOutOfBound
	}

	for i := index; i+1 < sa.length; i++ {
		sa.store[i] = sa.store[i+1]
	}
	sa.length--
	return nil
}
