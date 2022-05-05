package ll

import "errors"

var ErrorEmptyList = errors.New("list is empty")

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T any] interface {
	PushFront(data T)
	PopFront() (T, error)
	TopFront() (T, error)

	PushBack(data T)
	PopBack() (T, error)
	TopBack() (T, error)
	Empty() bool
}
