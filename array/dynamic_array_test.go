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

	if got.capacity != length {
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
