package sort

import (
	"errors"
)

type MinHeap []int

func NewMinHeap() MinHeap {
	return make([]int, 0)
}

func (heap *MinHeap) Empty() bool {
	return len(*heap) == 0
}

func (heap *MinHeap) Insert(value int) {
	*heap = append(*heap, value)

	(*heap).bubbleUp(len(*heap) - 1)
}

func (heap *MinHeap) bubbleUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := index / 2
	if (*heap)[index] < (*heap)[parentIndex] {
		(*heap)[index], (*heap)[parentIndex] = (*heap)[parentIndex], (*heap)[index]
	}

	(*heap).bubbleUp(parentIndex)
}

func (heap MinHeap) GetMin(value int) int {
	return heap[len(heap)-1]
}

func (heap *MinHeap) ExtractMin() (int, error) {
	if (*heap).Empty() {
		return 0, errors.New("cannot extract item from an empty heap")
	}

	if len(*heap) == 1 {
		res := (*heap)[0]
		(*heap) = NewMinHeap()
		return res, nil
	}

	first, end := 0, len(*heap)-1
	(*heap).swap(first, end)
	*heap = (*heap)[:end]
	(*heap).bubbleDown(first)
	return (*heap)[first], nil
}

func (heap *MinHeap) swap(index1, index2 int) {
	(*heap)[index1], (*heap)[index2] = (*heap)[index2], (*heap)[index1]
}

func (heap *MinHeap) bubbleDown(index int) {
	end := len(*heap) - 1

	leftIndex, rightIndex := 2*index+1, 2*index+2
	if leftIndex <= end && (*heap)[index] > (*heap)[leftIndex] {
		(*heap).swap(index, leftIndex)
		heap.bubbleDown(leftIndex)
	} else if rightIndex <= end && (*heap)[index] > (*heap)[rightIndex] {
		(*heap).swap(index, leftIndex)
		heap.bubbleDown(rightIndex)
	} else {
		return
	}
}

func HeapSort(s []int) {
	heap := NewMinHeap()
	for _, v := range s {
		heap.Insert(v)
	}

	for i := range s {
		s[i], _ = heap.ExtractMin()

	}
}
