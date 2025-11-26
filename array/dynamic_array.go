package array

// The Downside of StaticArray is that its size is fixed.
// To Solve this issue, we introduce dynamic array.

type DynamicArray struct {
	store    []int
	length   uint
	capacity uint
}

func NewDynamicArray(length uint) *DynamicArray {
	cap := length * 2
	return &DynamicArray{
		store:    make([]int, cap),
		length:   length,
		capacity: cap,
	}
}

func (da *DynamicArray) Read(index uint) (int, error) {
	if index >= da.capacity {
		return 0, ErrOutOfBound
	}

	return da.store[index], nil
}

func (da *DynamicArray) Update(index uint, value int) error {
	if index >= da.capacity {
		return ErrOutOfBound
	}

	da.store[index] = value
	return nil
}

// Mostly O(1)
// Worst case O(n)
func (da *DynamicArray) Insert(value int) {
	// when store is full
	if da.length == da.capacity {
		newStore := make([]int, da.capacity*2)
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
