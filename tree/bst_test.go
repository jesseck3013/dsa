package tree

import (
	"reflect"
	"testing"
)

func assertSlice(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestBST(t *testing.T) {
	root := NewNode(5)

	for i := range 10 {
		if i != 5 {
			node := NewNode(i)
			root.Insert(node)
		}
	}

	assertSlice(t, root.ToSlice(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestSearch(t *testing.T) {
	root := NewNode(5)
	dict := make(map[int]*BST)
	dict[5] = root

	for i := range 10 {
		if i != 5 {
			node := NewNode(i)
			root.Insert(node)
			dict[i] = node
		}
	}

	for k, v := range dict {
		if !root.Search(v) {
			t.Errorf("expcted node %d is in the tree, but it is not", k)
		}
	}

	node100 := NewNode(100)
	if root.Search(node100) {
		t.Error("expcted node 100 is not in the tree, but it is")
	}
}

func TestDelete(t *testing.T) {
	t.Run("Delete a node with no child", func(t *testing.T) {
		root := NewNode(5)
		dict := make(map[int]*BST)
		dict[5] = root

		for i := range 10 {
			if i != 5 {
				node := NewNode(i)
				root.Insert(node)
				dict[i] = node
			}
		}

		got := root.Delete(dict[4])
		assertSlice(t, got.ToSlice(), []int{0, 1, 2, 3, 5, 6, 7, 8, 9})
		got = root.Delete(dict[9])
		assertSlice(t, got.ToSlice(), []int{0, 1, 2, 3, 5, 6, 7, 8})
	})

	t.Run("Delete a node with one child", func(t *testing.T) {
		root := NewNode(5)
		dict := make(map[int]*BST)
		dict[5] = root

		for i := range 10 {
			if i != 5 {
				node := NewNode(i)
				root.Insert(node)
				dict[i] = node
			}
		}

		got := root.Delete(dict[3])
		assertSlice(t, got.ToSlice(), []int{0, 1, 2, 4, 5, 6, 7, 8, 9})
		got = root.Delete(dict[8])
		assertSlice(t, got.ToSlice(), []int{0, 1, 2, 4, 5, 6, 7, 9})
	})

	t.Run("Delete a node with two children", func(t *testing.T) {
		root := NewNode(5)
		dict := make(map[int]*BST)
		dict[5] = root

		for i := range 10 {
			if i != 5 {
				node := NewNode(i)
				root.Insert(node)
				dict[i] = node
			}
		}

		got := root.Delete(dict[5])
		assertSlice(t, got.ToSlice(), []int{0, 1, 2, 3, 4, 6, 7, 8, 9})
	})
}
