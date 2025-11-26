package list

import "testing"

func TestList(t *testing.T) {
	mt := NewMT()
	n3 := NewNode(3, mt)
	n2 := NewNode(2, n3)
	n1 := NewNode(1, n2)

	want := 3
	got := n1.Count()
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestListSum(t *testing.T) {
	mt := NewMT()
	n3 := NewNode(3, mt)
	n2 := NewNode(2, n3)
	n1 := NewNode(1, n2)

	want := 6
	got := n1.Sum()
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func biggerThan1(v int) bool {
	return v > 1
}

func TestFilter(t *testing.T) {
	mt := NewMT()
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
