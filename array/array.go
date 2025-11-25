package array

import "errors"

type StaticArray struct {
	store  []int
	length uint
}

func NewStaticArray(length uint) *StaticArray {
	store := make([]int, length)
	return &StaticArray{
		store:  store,
		length: length,
	}
}

func (sa *StaticArray) Read(index uint) (int, error) {
	if index >= sa.length {
		return 0, errors.New("index is out of bound")
	}

	return sa.store[index], nil
}
