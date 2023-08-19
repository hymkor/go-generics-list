package list

type Element[T any] struct {
	next  *Element[T]
	prev  *Element[T]
	Value T
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

type List[T any] struct {
	front *Element[T]
	back  *Element[T]
	count int
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Back() *Element[T] {
	return l.back
}

func (l *List[T]) Front() *Element[T] {
	return l.front
}

func (l *List[T]) Init() *List[T] {
	l.front = nil
	l.back = nil
	return l
}

func (l *List[T]) insertAfter(e, mark *Element[T]) {
	e.prev = mark
	e.next = mark.next

	if mark.next != nil {
		mark.next.prev = e
	} else {
		l.back = e
	}
	mark.next = e
}

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	e := &Element[T]{Value: v}
	l.insertAfter(e, mark)
	l.count++
	return e
}

func (l *List[T]) insertBefore(e, mark *Element[T]) {
	e.prev = mark.prev
	e.next = mark

	if mark.prev != nil {
		mark.prev.next = e
	} else {
		l.front = e
	}
	mark.prev = e
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	e := &Element[T]{Value: v}
	l.insertBefore(e, mark)
	l.count++
	return e
}

func (l *List[T]) Len() int {
	return l.count
}

func (l *List[T]) remove(e *Element[T]) {
	if e.prev != nil {
		e.prev.next = e.next
	} else {
		l.front = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	} else {
		l.back = e.prev
	}

	e.prev = nil
	e.next = nil
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e != mark {
		l.remove(e)
		l.insertAfter(e, mark)
	}
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.remove(e)
	l.insertBefore(e, mark)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.next != nil {
		l.remove(e)
		l.insertAfter(e, l.back)
	}
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.prev != nil {
		l.remove(e)
		l.insertBefore(e, l.front)
	}
}

func (l *List[T]) PushBack(v T) *Element[T] {
	e := &Element[T]{Value: v}
	if l.back == nil {
		l.back = e
		l.front = e
	} else {
		l.insertAfter(e, l.back)
	}
	l.count++
	return e
}

func (l *List[T]) PushBackList(other *List[T]) {
	for p := other.Front(); p != nil; p = p.Next() {
		l.PushBack(p.Value)
	}
}

func (l *List[T]) PushFront(v T) *Element[T] {
	e := &Element[T]{Value: v}
	if l.front == nil {
		l.front = e
		l.back = e
	} else {
		l.insertBefore(e, l.front)
	}
	l.count++
	return e
}

func (l *List[T]) PushFrontList(other *List[T]) {
	for p := other.Back(); p != nil; p = p.Prev() {
		l.PushFront(p.Value)
	}
}

func (l *List[T]) Remove(e *Element[T]) T {
	v := e.Value
	l.remove(e)
	l.count--
	return v
}
