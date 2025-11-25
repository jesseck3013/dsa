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

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		panic("expected no err")
	}
}

func AssertError(t *testing.T, want, got error) {
	if want != got {
		t.Errorf("expected error: %v, got %v", want, got)
	}
}

func AssertValue(t *testing.T, want, got int) {
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestSARead(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		want := 10
		sa.store[1] = want
		got, err := sa.Read(1)

		AssertNoError(t, err)
		AssertValue(t, got, want)
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		want := 10
		sa.store[1] = want
		_, err := sa.Read(100)

		AssertError(t, ErrOutOfBound, err)
	})
}

func TestSAUpdate(t *testing.T) {
	t.Run("In Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		want := 10
		sa.store[1] = want
		err := sa.Update(1, 10)

		AssertNoError(t, err)

		got, err := sa.Read(1)

		AssertNoError(t, err)
		AssertValue(t, want, got)
	})

	t.Run("out of Bound", func(t *testing.T) {
		sa := NewStaticArray(10)
		err := sa.Update(100, 10)

		AssertError(t, ErrOutOfBound, err)
	})
}
