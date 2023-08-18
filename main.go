package list

import (
	"container/list"
)

type Element[T any] struct {
	value *list.Element
}

func (e Element[T]) IsSome() bool {
	return e.value != nil
}

func (e Element[T]) IsNone() bool {
	return e.value == nil
}

func (e Element[T]) Value() T {
	return e.value.Value.(T)
}

func (e Element[T]) Next() Element[T] {
	return Element[T]{value: e.value.Next()}
}

func (e Element[T]) Prev() Element[T] {
	return Element[T]{value: e.value.Prev()}
}

type List[T any] struct {
	list *list.List
}

func New[T any]() List[T] {
	return List[T]{list: list.New()}
}

func (l List[T]) Back() Element[T] {
	return Element[T]{value: l.list.Back()}
}

func (l List[T]) Front() Element[T] {
	return Element[T]{value: l.list.Front()}
}

func (l List[T]) Init() List[T] {
	return List[T]{list: l.list.Init()}
}

func (l List[T]) InsertAfter(v T, mark Element[T]) Element[T] {
	return Element[T]{value: l.list.InsertAfter(v, mark.value)}
}

func (l List[T]) InsertBefore(v T, mark Element[T]) Element[T] {
	return Element[T]{value: l.list.InsertBefore(v, mark.value)}
}

func (l List[T]) Len() int {
	return l.list.Len()
}

func (l List[T]) MoveAfter(e, mark Element[T]) {
	l.list.MoveAfter(e.value, mark.value)
}

func (l List[T]) MoveBefore(e, mark Element[T]) {
	l.list.MoveBefore(e.value, mark.value)
}

func (l List[T]) MoveToBack(e Element[T]) {
	l.list.MoveToBack(e.value)
}

func (l List[T]) MoveToFront(e Element[T]) {
	l.list.MoveToFront(e.value)
}

func (l List[T]) PushBack(v T) Element[T] {
	return Element[T]{value: l.list.PushBack(v)}
}

func (l List[T]) PushBackList(other List[T]) {
	l.list.PushBackList(other.list)
}

func (l List[T]) PushFront(v T) Element[T] {
	return Element[T]{value: l.list.PushFront(v)}
}

func (l List[T]) PushFrontList(other List[T]) {
	l.list.PushFrontList(other.list)
}

func (l List[T]) Remove(e Element[T]) T {
	return l.list.Remove(e.value).(T)
}
