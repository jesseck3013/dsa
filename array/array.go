package array

type StaticArray struct {
	store  []int
	length uint
}

type StaticArrayErr string

func (err StaticArrayErr) Error() string {
	return string(err)
}

const ErrOutOfBound = StaticArrayErr("index is out of bound")

// Time: O(n)
// Space: O(n)
func NewStaticArray(length uint) *StaticArray {
	store := make([]int, length)
	return &StaticArray{
		store:  store,
		length: length,
	}
}

// Time: O(1)
// Space: O(1)
func (sa *StaticArray) Read(index uint) (int, error) {
	if index >= sa.length {
		return 0, ErrOutOfBound
	}

	return sa.store[index], nil
}

// Time: O(1)
// Space: O(1)
func (sa *StaticArray) Update(index uint, value int) error {
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
func (sa *StaticArray) Delete(index uint) error {
	if index >= sa.length {
		return ErrOutOfBound
	}

	for i := index; i+1 < sa.length; i++ {
		sa.store[i] = sa.store[i+1]
	}
	sa.length--
	return nil
}
