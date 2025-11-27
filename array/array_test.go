package array

import (
	"testing"

	testutils "github.com/jesseck3013/dsa/test_utils"
)

func TestStaticArray(t *testing.T) {
	length := 100
	got := NewStaticArray[int](100)

	for i := range length {
		if got.store[i] != 0 {
			t.Errorf("expected %d, got %d, at index %d", 0, got.store[i], i)
		}
	}
}

func TestSARead(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		want := 10
		sa.store[1] = want
		got, err := sa.Read(1)

		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		_, err := sa.Read(100)
		testutils.AssertError(t, ErrOutOfBound, err)
	})
}

func TestSAUpdate(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		want := 10
		sa.store[1] = want
		err := sa.Update(1, 10)

		testutils.AssertNoError(t, err)

		got, err := sa.Read(1)

		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, want, got)
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		err := sa.Update(100, 10)

		testutils.AssertError(t, ErrOutOfBound, err)
	})
}

func TestSADelete(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		sa.Update(3, 10)
		sa.Delete(3)

		if sa.length != 9 {
			t.Errorf("expected length %d, got %d", 9, sa.length)
		}

		for i := range 9 {
			if sa.store[i] != 0 {
				t.Errorf("expected value %d, got %d, at index %d", 0, sa.store[i], i)
			}
		}
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray[int](10)
		err := sa.Delete(100)

		testutils.AssertError(t, ErrOutOfBound, err)
	})
}
