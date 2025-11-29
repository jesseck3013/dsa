package list

type LinkedList[T any] struct {
	Head  Lister[T]
	Tail  Lister[T]
	empty bool
	len   int
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

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Exist(n *Element[T]) bool {
	if l.Head.Same(n) {
		return true
	}
	return false
}

func (l *LinkedList[T]) Delete(n *Element[T]) {
	if l.len == 0 {
		return
	}

	if l.Head.Exist(n) {
		prev := n.prev
		next := n.next
		prev.SetNext(next)
		next.SetPrev(prev)

		if l.Head == n {
			l.Head = n.next
		}

		if l.Tail == n {
			l.Tail = n.prev
		}
		n.next = nil // for gc to work
		n.prev = nil // for gc to work
		l.len--
		if l.len == 0 {
			l.empty = true
		}
	}
}

func (l *LinkedList[T]) InsertHead(n *Element[T]) {
	newHead := l.Head.InsertPrev(n)
	l.Head = newHead
	l.len++
	if l.Empty() {
		l.Tail = newHead
		l.empty = false
	}
}

func (l *LinkedList[T]) InsertTail(n *Element[T]) {
	newTail := l.Tail.InsertNext(n)
	l.Tail = newTail
	l.len++
	if l.Empty() {
		l.Head = newTail
		l.empty = false
	}
}

type Lister[T any] interface {
	InsertPrev(*Element[T]) Lister[T]
	InsertNext(*Element[T]) Lister[T]

	SetPrev(Lister[T])
	SetNext(Lister[T])

	ToSlice() []T
	ToSliceRev() []T

	Exist(*Element[T]) bool

	Same(Lister[T]) bool
	sameMT(*EmptyElement[T]) bool
	sameElement(*Element[T]) bool
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

func (mt *EmptyElement[T]) Same(elem Lister[T]) bool {
	return elem.sameMT(mt)
}

func (mt *EmptyElement[T]) sameMT(elem *EmptyElement[T]) bool {
	return mt == elem
}

func (mt *EmptyElement[T]) sameElement(elem *Element[T]) bool {
	return false
}

func (mt *EmptyElement[T]) Exist(elem *Element[T]) bool {
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

func (e *Element[T]) ToSlice() []T {
	return append([]T{e.value}, e.next.ToSlice()...)
}

func (e *Element[T]) ToSliceRev() []T {
	return append([]T{e.value}, e.prev.ToSliceRev()...)
}

func (e *Element[T]) Same(elem Lister[T]) bool {
	return elem.sameElement(e)
}

func (e *Element[T]) sameMT(elem *EmptyElement[T]) bool {
	return false
}

func (e *Element[T]) sameElement(elem *Element[T]) bool {
	return e == elem
}

func (e *Element[T]) Exist(elem *Element[T]) bool {
	if e.Same(elem) {
		return true
	}

	return e.next.Exist(elem)
}
