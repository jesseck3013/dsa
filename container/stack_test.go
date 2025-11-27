package container

import (
	"testing"

	testutils "github.com/jesseck3013/dsa/test_utils"
)

func TestStack(t *testing.T) {
	t.Run("non-empty stack Top", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		want := 3
		got, err := s.Top()
		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, want, got)
	})

	t.Run("empty stack Top", func(t *testing.T) {
		s := NewStack[int]()

		_, err := s.Top()
		testutils.AssertError(t, ErrEmpty, err)
	})
}

func TestPop(t *testing.T) {
	t.Run("non-empty Pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		want := 3
		got, err := s.Pop()
		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, want, got)

		wantTop := 2
		gotTop, err := s.Top()
		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, wantTop, gotTop)
	})

	t.Run("empty stack Pop", func(t *testing.T) {
		s := NewStack[int]()

		_, err := s.Pop()
		testutils.AssertError(t, ErrEmpty, err)
	})
}
