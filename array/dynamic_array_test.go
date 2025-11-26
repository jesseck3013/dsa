package array

import (
	"testing"
)

func TestDynamicArray(t *testing.T) {
	var length uint = 100
	got := NewDynamicArray(100)

	if got.length != length {
		t.Errorf("expected length %d, got %d", length, got.length)
	}

	if got.capacity != length*2 {
		t.Errorf("expected capacity %d, got %d", length, got.capacity)
	}

	for i := range length {
		if got.store[i] != 0 {
			t.Errorf("expected %d, got %d, at index %d", 0, got.store[i], i)
		}
	}
}

func TestDARead(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		da := NewDynamicArray(10)
		want := 10
		da.store[1] = want
		got, err := da.Read(1)

		AssertNoError(t, err)
		AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		da := NewDynamicArray(10)
		_, err := da.Read(100)
		AssertError(t, ErrOutOfBound, err)
	})
}

func TestDAUpdate(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		da := NewDynamicArray(10)
		want := 10
		err := da.Update(2, want)
		AssertNoError(t, err)

		got, err := da.Read(2)
		AssertNoError(t, err)
		AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		da := NewDynamicArray(10)
		err := da.Update(100, 100)
		AssertError(t, ErrOutOfBound, err)
	})
}

func TestDAInsert(t *testing.T) {
	t.Run("store is not full", func(t *testing.T) {
		da := NewDynamicArray(10)
		want := 99
		da.Insert(want)

		var wantLength uint = 11
		if da.length != wantLength {
			t.Errorf("expected length %d, got %d", wantLength, da.length)
		}

		var wantCapacity uint = 20
		if da.capacity != wantCapacity {
			t.Errorf("expected capacity %d, got %d", wantCapacity, da.length)
		}

		got, err := da.Read(10)
		AssertNoError(t, err)
		AssertValue(t, want, got)
	})

	t.Run("store is full", func(t *testing.T) {
		da := NewDynamicArray(1)
		da.Insert(10)
		// da is full

		want := 99
		da.Insert(want)

		var wantLength uint = 3
		if da.length != 3 {
			t.Errorf("expected length %d, got %d", wantLength, da.length)
		}

		var wantCapacity uint = 4
		if da.capacity != wantCapacity {
			t.Errorf("expected capacity %d, got %d", wantCapacity, da.length)
		}

		got, err := da.Read(2)
		AssertNoError(t, err)
		AssertValue(t, want, got)
	})
}
