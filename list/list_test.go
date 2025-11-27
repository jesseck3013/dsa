package list

import "testing"

func TestList(t *testing.T) {
	mt := NewMT[int]()
	n3 := NewNode(3, mt)
	n2 := NewNode(2, n3)
	n1 := NewNode(1, n2)

	want := 3
	got := n1.Count()
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func biggerThan1(v int) bool {
	return v > 1
}

func TestFilter(t *testing.T) {
	mt := NewMT[int]()
	n3 := NewNode(3, mt)
	n2 := NewNode(2, n3)
	n1 := NewNode(1, n2)

	got := n1.Filter(biggerThan1)

	wantCount := 2
	gotCount := got.Count()
	if gotCount != wantCount {
		t.Errorf("expected count %d, got %d", wantCount, gotCount)
	}
}

func TestDelete(t *testing.T) {
	mt := NewMT[int]()
	n3 := NewNode(3, mt)
	n2 := NewNode(2, n3)
	n1 := NewNode(1, n2)

	n := n1.Delete(n2)

	wantCount := 2
	gotCount := n.Count()
	if gotCount != wantCount {
		t.Errorf("expected count %d, got %d", wantCount, gotCount)
	}
}

func TestSame(t *testing.T) {
	t.Run("two empty lists are same", func(t *testing.T) {
		mt1 := NewMT[int]()
		mt2 := NewMT[int]()
		got := mt1.Same(mt2)
		want := true
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

	t.Run("two non-empty lists are same", func(t *testing.T) {
		mt1 := NewMT[int]()
		n1 := NewNode(1, mt1)

		mt2 := NewMT[int]()
		n2 := NewNode(1, mt2)

		got := n1.Same(n2)
		want := true
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}

func TestSearch(t *testing.T) {
	t.Run("Search a node in the list", func(t *testing.T) {
		mt := NewMT[int]()
		n3 := NewNode(3, mt)
		n2 := NewNode(2, n3)
		n1 := NewNode(1, n2)

		want := 1
		got := n1.Search(n2)
		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("Search a node out of the list", func(t *testing.T) {
		mt := NewMT[int]()
		n4 := NewNode(4, mt)
		n3 := NewNode(3, mt)
		n2 := NewNode(2, n3)
		n1 := NewNode(1, n2)

		want := -1
		got := n1.Search(n4)
		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
