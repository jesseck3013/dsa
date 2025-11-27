package container

import (
	"testing"

	testutils "github.com/jesseck3013/dsa/test_utils"
)

func TestQueue(t *testing.T) {
	t.Run("non-empty dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		for i := range 3 {
			q.Enqueue(i + 1)
		}

		want := 1
		got, err := q.Dequeue()
		testutils.AssertNoError(t, err)
		testutils.AssertValue(t, want, got)
	})

	t.Run("empty queue dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		_, err := q.Dequeue()
		testutils.AssertError(t, ErrQueueEmpty, err)
	})
}
