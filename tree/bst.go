package tree

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func NewBST(value int, left, right *BST) *BST {
	return &BST{
		Value: value,
		Left:  left,
		Right: right,
	}
}

func NewNode(value int) *BST {
	return NewBST(value, nil, nil)
}

func (b *BST) Insert(node *BST) {
	if node != nil {
		switch {
		case node.Value > b.Value:
			if b.Right == nil {
				b.Right = node
			} else {
				b.Right.Insert(node)
			}
		case node.Value < b.Value:
			if b.Left == nil {
				b.Left = node
			} else {
				b.Left.Insert(node)
			}
		}
	}
}

func (b *BST) Search(node *BST) bool {
	if b == nil {
		return false
	}

	if node == nil {
		return false
	}

	if node.Value > b.Value {
		return b.Right.Search(node)
	}

	if node.Value < b.Value {
		return b.Left.Search(node)
	}

	if b != node {
		return b.Left.Search(node) || b.Right.Search(node)
	}

	return true
}

func (b *BST) Delete(node *BST) *BST {
	if b == node {
		return processRemovedNode(b)
	}

	deleteHelper(b, node)
	return b
}

func processRemovedNode(node *BST) *BST {
	if node == nil {
		return node
	}

	if node.Left == nil && node.Right == nil {
		return nil
	}

	if node.Left == nil {
		return node.Right
	}

	if node.Right == nil {
		return node.Left
	}

	newRoot := node.Right.FindMin()
	newRight := node.Right.Delete(newRoot)
	newRoot.Left = node.Left
	newRoot.Right = newRight
	node.Left = nil
	node.Right = nil
	return newRoot
}

func (b *BST) FindMin() *BST {
	min := b

	for min.Left != nil {
		min = min.Left
	}
	return min
}

func deleteHelper(root *BST, node *BST) {
	if root == nil {
		return
	}

	if root.Left == node {
		root.Left = processRemovedNode(node)
	}

	if root.Right == node {
		root.Right = processRemovedNode(node)
	}

	deleteHelper(root.Left, node)
	deleteHelper(root.Right, node)
}

func (root *BST) ToSlice() []int {
	return inOrder(root, []int{})
}

func inOrder(root *BST, acc []int) []int {
	if root == nil {
		return acc
	}

	acc = inOrder(root.Left, acc)
	acc = append(acc, root.Value)
	return inOrder(root.Right, acc)
}
