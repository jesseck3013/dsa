package array

// The Downside of StaticArray is that its size is fixed.
// To Solve this issue, we introduce dynamic array.

type DynamicArray[T any] struct {
	store    []T
	length   uint
	capacity uint
}

func NewDynamicArray[T any](length uint) *DynamicArray[T] {
	cap := length * 2
	return &DynamicArray[T]{
		store:    make([]T, cap),
		length:   length,
		capacity: cap,
	}
}

func (da *DynamicArray[T]) Read(index uint) (T, error) {
	if index >= da.capacity {
		var zero T
		return zero, ErrOutOfBound
	}

	return da.store[index], nil
}

func (da *DynamicArray[T]) Update(index uint, value T) error {
	if index >= da.capacity {
		return ErrOutOfBound
	}

	da.store[index] = value
	return nil
}

// Mostly O(1)
// Worst case O(n)
func (da *DynamicArray[T]) Insert(value T) {
	// when store is full
	if da.length == da.capacity {
		newStore := make([]T, da.capacity*2)
		copy(newStore, da.store)
		newStore[da.length] = value
		da.store = newStore
		da.length++
		da.capacity *= 2
	} else {
		da.store[da.length] = value
		da.length++
	}
}

func shiftLeft[T any](s []T, index uint) {
	for i := index; i < uint(len(s)-1); i++ {
		s[i] = s[i+1]
	}
}

func (da *DynamicArray[T]) Delete(index uint) error {
	if index < da.capacity {
		shiftLeft(da.store, index)
		da.length--
		return nil
	} else {
		return ErrOutOfBound
	}
}
