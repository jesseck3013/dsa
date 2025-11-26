package array

// The Downside of StaticArray is that its size is fixed.
// To Solve this issue, we introduce dynamic array.

type DynamicArray struct {
	store    []int
	length   uint
	capacity uint
}

func NewDynamicArray(length uint) *DynamicArray {
	return &DynamicArray{
		store:    make([]int, length),
		length:   length,
		capacity: length,
	}
}

func (da *DynamicArray) Read(index uint) (int, error) {
	if index >= da.capacity {
		return 0, ErrOutOfBound
	}

	return da.store[index], nil
}
