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

func TestDoublyList(t *testing.T) {
	node1 := NewElement(1)
	node2 := NewElement(2)
	node3 := NewElement(3)
	node4 := NewElement(4)
	node5 := NewElement(5)

	linkedList := NewLinkedList[int]()

	linkedList.InsertTail(node1)
	linkedList.InsertTail(node2)
	linkedList.InsertTail(node3)
	linkedList.InsertTail(node4)
	node2.InsertNext(node5)

	AssertList(t, linkedList, []int{1, 2, 5, 3, 4})
}
