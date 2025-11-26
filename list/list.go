package list

type CompareFunc func(int) bool

type List interface {
	Count() int
	Sum() int
	Filter(CompareFunc) List
	Delete(*Node) List

	Same(List) bool
	SameNode(*Node) bool
	SameMT(*MT) bool
}

// Represent an empty list
type MT struct {
}

func NewMT() *MT {
	return &MT{}
}

func (mt *MT) Count() int {
	return 0
}

func (mt *MT) Sum() int {
	return 0
}

func (mt *MT) Filter(comp CompareFunc) List {
	return mt
}

func (mt *MT) Delete(target *Node) List {
	return mt
}

func (mt *MT) Same(l List) bool {
	return l.SameMT(mt)
}

func (mt *MT) SameNode(n *Node) bool {
	return false
}

func (mt *MT) SameMT(n *MT) bool {
	return true
}

type Node struct {
	value int
	rest  List
}

func (n *Node) Count() int {
	return 1 + n.rest.Count()
}

func (n *Node) Sum() int {
	return n.value + n.rest.Sum()
}

func (n *Node) Filter(comp CompareFunc) List {
	if comp(n.value) {
		return NewNode(n.value, n.rest.Filter(comp))
	}
	return n.rest.Filter(comp)
}

func (n *Node) Delete(target *Node) List {
	if n == target {
		return n.rest
	}

	return NewNode(n.value, n.rest.Delete(target))
}

func (n *Node) Same(l List) bool {
	return l.SameNode(n)
}

func (n *Node) SameValue(v int) bool {
	return n.value == v
}

func (n *Node) SameNode(target *Node) bool {
	if n.value == target.value {
		return n.rest.Same(target.rest)
	}

	return false
}

func (n *Node) SameMT(mt *MT) bool {
	return false
}

func NewNode(value int, rest List) *Node {
	return &Node{
		value: value,
		rest:  rest,
	}
}
