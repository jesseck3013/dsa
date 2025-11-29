package list

type LinkedList[T any] struct {
	Head  Lister[T]
	Tail  Lister[T]
	empty bool

	next Lister[T]
	prev Lister[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head:  NewDoublyMT[T](),
		Tail:  NewDoublyMT[T](),
		empty: true,
	}
}

func (l *LinkedList[T]) Empty() bool {
	return l.empty
}

func (l *LinkedList[T]) InsertHead(n *Element[T]) {
	newHead := l.Head.InsertPrev(n)
	l.Head = newHead
	if l.Empty() {
		l.Tail = newHead
		l.empty = false
	}
}

func (l *LinkedList[T]) InsertTail(n *Element[T]) {
	newTail := l.Tail.InsertNext(n)
	l.Tail = newTail
	if l.Empty() {
		l.Head = newTail
		l.empty = false
	}
}

func (l *LinkedList[T]) HasNext() bool {
	return l.next.HasNext()
}

func (l *LinkedList[T]) HasPrev() bool {
	return l.next.HasPrev()
}

type Lister[T any] interface {
	InsertPrev(*Element[T]) Lister[T]
	InsertNext(*Element[T]) Lister[T]

	SetPrev(Lister[T])
	SetNext(Lister[T])

	ToSlice() []T
	ToSliceRev() []T

	HasNext() bool
	HasPrev() bool
}

type EmptyElement[T any] struct {
}

func NewDoublyMT[T any]() *EmptyElement[T] {
	return &EmptyElement[T]{}
}

func (mt *EmptyElement[T]) InsertPrev(newPrev *Element[T]) Lister[T] {
	newPrev.SetNext(mt)
	return newPrev
}

func (mt *EmptyElement[T]) InsertNext(newNext *Element[T]) Lister[T] {
	newNext.SetPrev(mt)
	return newNext
}

func (mt *EmptyElement[T]) SetPrev(elem Lister[T]) {
	return
}

func (mt *EmptyElement[T]) SetNext(elem Lister[T]) {
	return
}

func (mt *EmptyElement[T]) HasNext() bool {
	return false
}

func (mt *EmptyElement[T]) HasPrev() bool {
	return false
}

func (mt *EmptyElement[T]) ToSlice() []T {
	return []T{}
}

func (mt *EmptyElement[T]) ToSliceRev() []T {
	return []T{}
}

type Element[T any] struct {
	value T
	prev  Lister[T]
	next  Lister[T]
}

func NewElement[T any](value T) *Element[T] {
	return &Element[T]{
		value: value,
		prev:  NewDoublyMT[T](),
		next:  NewDoublyMT[T](),
	}
}

func (n *Element[T]) SetPrev(prev Lister[T]) {
	n.prev = prev
}

func (n *Element[T]) SetNext(next Lister[T]) {
	n.next = next
}

func (n *Element[T]) InsertPrev(newPrev *Element[T]) Lister[T] {
	oldPrev := n.prev
	newPrev.SetNext(n)
	n.SetPrev(newPrev)

	oldPrev.SetNext(newPrev)
	newPrev.SetPrev(oldPrev)

	return newPrev
}

func (n *Element[T]) InsertNext(newNext *Element[T]) Lister[T] {
	oldNext := n.next
	n.SetNext(newNext)
	newNext.SetPrev(n)

	oldNext.SetPrev(newNext)
	newNext.SetNext(oldNext)

	return newNext
}

func (n *Element[T]) HasNext() bool {
	return true
}

func (n *Element[T]) HasPrev() bool {
	return true
}

func (mt *Element[T]) ToSlice() []T {
	return append([]T{mt.value}, mt.next.ToSlice()...)
}

func (mt *Element[T]) ToSliceRev() []T {
	return append([]T{mt.value}, mt.prev.ToSliceRev()...)
}
