package sort

import "testing"

func assertSorted(t *testing.T, s []int, wantLength int) {
	t.Helper()

	if len(s) != wantLength {
		t.Errorf("expcted slice length %v, got %v", wantLength, len(s))
	}

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

type TestCase struct {
	Name           string
	Unsorted       []int
	ExpectedLength int
}

var testCases = []TestCase{
	TestCase{"sort unsorted", []int{3, 2, 4, 6, 10}, 5},
	TestCase{"sort empty", []int{}, 0},
	TestCase{"sort 1-item slice", []int{3}, 1},
	TestCase{"sort sorted", []int{1, 2, 3, 4, 5}, 5},
}

func TestHeapSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			HeapSort(testCase.Unsorted)
			assertSorted(t, testCase.Unsorted, testCase.ExpectedLength)
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			MergeSort(testCase.Unsorted)
			assertSorted(t, testCase.Unsorted, testCase.ExpectedLength)
		})
	}
}
