package list

type CompareFunc func(int) bool

type List interface {
	Count() int
	Sum() int
	Filter(CompareFunc) List
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

func NewNode(value int, rest List) *Node {
	return &Node{
		value: value,
		rest:  rest,
	}
}
