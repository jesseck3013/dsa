package array

import (
	"testing"
)

func TestStaticArray(t *testing.T) {
	length := 100
	got := NewStaticArray(100)

	for i := range length {
		if got.store[i] != 0 {
			t.Errorf("expected %d, got %d, at index %d", 0, got.store[i], i)
		}
	}
}

func TestSARead(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		want := 10
		sa.store[1] = want
		got, err := sa.Read(1)

		if err != nil {
			t.Error(err)
			panic("not expected err")
		}

		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		want := 10
		sa.store[1] = want
		_, err := sa.Read(100)

		if err == nil {
			t.Error(err)
			panic("expected out of bound error")
		}
	})
}
