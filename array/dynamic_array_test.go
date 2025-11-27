package array

import (
	"testing"
)

func AssertLenCap[T any](t *testing.T, da *DynamicArray[T], wantLength, wantCapacity uint) {
	if da.length != wantLength {
		t.Errorf("expected length %d, got %d", wantLength, da.length)
	}

	if da.capacity != wantCapacity {
		t.Errorf("expected capacity %d, got %d", wantCapacity, da.length)
	}
}

func TestDynamicArray(t *testing.T) {
	var length uint = 100
	got := NewDynamicArray[int](100)

	AssertLenCap(t, got, length, length*2)

	for i := range length {
		if got.store[i] != 0 {
			t.Errorf("expected %d, got %d, at index %d", 0, got.store[i], i)
		}
	}
}

func TestDARead(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		da := NewDynamicArray[int](10)
		want := 10
		da.store[1] = want
		got, err := da.Read(1)

		AssertNoError(t, err)
		AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		da := NewDynamicArray[int](10)
		_, err := da.Read(100)
		AssertError(t, ErrOutOfBound, err)
	})
}

func TestDAUpdate(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		da := NewDynamicArray[int](10)
		want := 10
		err := da.Update(2, want)
		AssertNoError(t, err)

		got, err := da.Read(2)
		AssertNoError(t, err)
		AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		da := NewDynamicArray[int](10)
		err := da.Update(100, 100)
		AssertError(t, ErrOutOfBound, err)
	})
}

func TestDAInsert(t *testing.T) {
	t.Run("store is not full", func(t *testing.T) {
		da := NewDynamicArray[int](10)
		want := 99
		da.Insert(want)

		AssertLenCap(t, da, 11, 20)

		got, err := da.Read(10)
		AssertNoError(t, err)
		AssertValue(t, want, got)
	})

	t.Run("store is full", func(t *testing.T) {
		da := NewDynamicArray[int](1)
		da.Insert(10)
		// da is full

		want := 99
		da.Insert(want)

		AssertLenCap(t, da, 3, 4)

		got, err := da.Read(2)
		AssertNoError(t, err)
		AssertValue(t, want, got)
	})
}

func TestDADelete(t *testing.T) {
	t.Run("in bound", func(t *testing.T) {
		da := NewDynamicArray[int](4)
		for i := range 4 {
			da.Update(uint(i), i)
		}

		err := da.Delete(0)
		AssertNoError(t, err)

		for i := range 3 {
			got, err := da.Read(uint(i))
			AssertNoError(t, err)
			AssertValue(t, i+1, got)
		}

		AssertLenCap(t, da, 3, 8)
	})

	t.Run("out of bound", func(t *testing.T) {
		da := NewDynamicArray[int](0)
		err := da.Delete(0)
		AssertError(t, ErrOutOfBound, err)
	})
}
