package list

type CompareFunc[T comparable] func(T) bool

type List[T comparable] interface {
	Count() int
	Filter(CompareFunc[T]) List[T]
	Delete(*Node[T]) List[T]

	Insert(*Node[T]) List[T]

	Same(List[T]) bool
	SameNode(*Node[T]) bool
	SameMT(*MT[T]) bool
	Search(*Node[T]) int
	ToSlice() []T
}

// Represent an empty list
type MT[T comparable] struct {
}

func NewMT[T comparable]() *MT[T] {
	return &MT[T]{}
}

func (mt *MT[T]) Count() int {
	return 0
}

func (mt *MT[T]) Filter(comp CompareFunc[T]) List[T] {
	return mt
}

func (mt *MT[T]) Delete(target *Node[T]) List[T] {
	return mt
}

func (mt *MT[T]) Insert(n *Node[T]) List[T] {
	n.rest = mt
	return n
}

func (mt *MT[T]) ToSlice() []T {
	return make([]T, 0)
}

func (mt *MT[T]) Same(l List[T]) bool {
	return l.SameMT(mt)
}

func (mt *MT[T]) SameNode(n *Node[T]) bool {
	return false
}

func (mt *MT[T]) SameMT(n *MT[T]) bool {
	return true
}

func (mt *MT[T]) Search(n *Node[T]) int {
	return -1
}

type Node[T comparable] struct {
	value T
	rest  List[T]
}

func (n *Node[T]) Count() int {
	return 1 + n.rest.Count()
}

func (n *Node[T]) Filter(comp CompareFunc[T]) List[T] {
	if comp(n.value) {
		return NewNode(n.value, n.rest.Filter(comp))
	}
	return n.rest.Filter(comp)
}

func (n *Node[T]) Delete(target *Node[T]) List[T] {
	if n == target {
		return n.rest
	}

	return NewNode(n.value, n.rest.Delete(target))
}

func (n *Node[T]) Same(l List[T]) bool {
	return l.SameNode(n)
}

func (n *Node[T]) SameValue(v T) bool {
	return n.value == v
}

func (n *Node[T]) SameNode(target *Node[T]) bool {
	if n.value == target.value {
		return n.rest.Same(target.rest)
	}

	return false
}

func (n *Node[T]) SameMT(mt *MT[T]) bool {
	return false
}

func (n *Node[T]) Search(target *Node[T]) int {
	if n == target {
		return 0
	}

	if n.rest.Search(target) == -1 {
		return -1
	}
	return 1 + n.rest.Search(target)
}

func (n *Node[T]) Insert(head *Node[T]) List[T] {
	head.rest = n
	return head
}

func (n *Node[T]) ToSlice() []T {
	return append([]T{n.value}, n.rest.ToSlice()...)
}

func NewNode[T comparable](value T, rest List[T]) *Node[T] {
	return &Node[T]{
		value: value,
		rest:  rest,
	}
}
