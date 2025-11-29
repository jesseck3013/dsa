package list

import (
	"slices"
	"testing"

	testutils "github.com/jesseck3013/dsa/test_utils"
)

func reverseSlice[T any](origin []T) []T {
	res := make([]T, len(origin))
	copy(res, origin)
	slices.Reverse(res)
	return res
}

func AssertLength[T any](t testing.TB, got, want []T) {
	expectedLength := len(want)
	gotLength := len(got)
	if gotLength != expectedLength {
		t.Errorf("expected length %d, got %d", expectedLength, gotLength)
	}
}

func AssertList[T comparable](t *testing.T, got *LinkedList[T], want []T) {
	t.Helper()

	headToTail := got.Head.ToSlice()

	AssertLength(t, headToTail, want)
	for i, got := range headToTail {
		testutils.AssertValue(t, want[i], got)
	}

	tailToHead := got.Tail.ToSliceRev()

	wantHeadToTail := reverseSlice(want)
	AssertLength(t, tailToHead, wantHeadToTail)
	for i, got := range tailToHead {
		testutils.AssertValue(t, wantHeadToTail[i], got)
	}
}

func newTestList[T any](s []*Element[T]) *LinkedList[T] {
	linkedList := NewLinkedList[T]()

	for _, v := range s {
		linkedList.InsertTail(v)
	}

	return linkedList
}

func TestDoublyList(t *testing.T) {
	node1 := NewElement(1)
	node2 := NewElement(2)
	node3 := NewElement(3)
	node4 := NewElement(4)
	node5 := NewElement(5)
	linkedList := newTestList([]*Element[int]{node1, node2, node5, node3, node4})
	AssertList(t, linkedList, []int{1, 2, 5, 3, 4})
}

func TestDoublyListDelete(t *testing.T) {
	t.Run("Delete head", func(t *testing.T) {
		node1 := NewElement(1)
		node2 := NewElement(2)
		node3 := NewElement(3)
		node4 := NewElement(4)
		node5 := NewElement(5)
		linkedList := newTestList([]*Element[int]{node1, node2, node5, node3, node4})
		linkedList.Delete(node1)
		if linkedList.Head != node2 {
			t.Errorf("expected head %v, got %v", node2, linkedList.Head)
		}
		AssertList(t, linkedList, []int{2, 5, 3, 4})
	})

	t.Run("Delete tail", func(t *testing.T) {
		node1 := NewElement(1)
		node2 := NewElement(2)
		node3 := NewElement(3)
		node4 := NewElement(4)
		node5 := NewElement(5)
		linkedList := newTestList([]*Element[int]{node1, node2, node5, node3, node4})
		linkedList.Delete(node4)
		if linkedList.Tail != node3 {
			t.Errorf("expected tail %v, got %v", node3, linkedList.Tail)
		}
		AssertList(t, linkedList, []int{1, 2, 5, 3})
	})

	t.Run("Delete mid", func(t *testing.T) {
		node1 := NewElement(1)
		node2 := NewElement(2)
		node3 := NewElement(3)
		node4 := NewElement(4)
		node5 := NewElement(5)
		linkedList := newTestList([]*Element[int]{node1, node2, node5, node3, node4})
		linkedList.Delete(node5)

		if linkedList.Head != node1 {
			t.Errorf("expected head %v, got %v", node1, linkedList.Head)
		}
		if linkedList.Tail != node4 {
			t.Errorf("expected tail %v, got %v", node4, linkedList.Tail)
		}
		AssertList(t, linkedList, []int{1, 2, 3, 4})
	})
}
