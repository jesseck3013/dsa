package sort

import "testing"

func assertSorted(t *testing.T, s []int) {
	t.Helper()
	if len(s) == 0 {
		return
	}

	cur := s[0]
	for _, v := range s {
		if cur > v {
			t.Errorf("the given slice is not sorted: %v", s)
			break
		}
	}
}

func TestHeap(t *testing.T) {
	t.Run("sort unsorted", func(t *testing.T) {
		unsorted := []int{3, 2, 4, 6, 10}
		HeapSort(unsorted)
		assertSorted(t, unsorted)
	})

	t.Run("sort empty", func(t *testing.T) {
		unsorted := []int{}
		HeapSort(unsorted)
		assertSorted(t, unsorted)
	})

	t.Run("sort sorted", func(t *testing.T) {
		unsorted := []int{1, 2, 3, 4, 5}
		HeapSort(unsorted)
		assertSorted(t, unsorted)
	})
}
