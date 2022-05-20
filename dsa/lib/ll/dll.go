package ll

type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (dll *DoublyLinkedList[T]) PushFront(data T) {
	n := &Node[T]{Data: data}
	if dll.Head == nil {
		dll.Head = n
		dll.Tail = n
		return
	}

	n.Next = dll.Head
	n.Next.Prev = n
	dll.Head = n
	if dll.Tail == n.Prev {
		dll.Tail = n
	}
}

func (dll *DoublyLinkedList[T]) PopFront() (T, error) {
	if dll.Head == nil {
		var z T
		return z, ErrorEmptyList
	}
	n := dll.Head
	dll.Head = n.Next

	if dll.Head == nil {
		dll.Tail = nil
	} else {
		dll.Head.Prev = nil
	}
	return n.Data, nil

}

func (dll *DoublyLinkedList[T]) TopFront() (T, error) {
	if dll.Head == nil {
		var z T
		return z, ErrorEmptyList
	}
	return dll.Head.Data, nil
}

func (dll *DoublyLinkedList[T]) PushBack(data T) {
	n := &Node[T]{Data: data}
	if dll.Tail == nil {
		dll.Head = n
		dll.Tail = n
		return
	}

	n.Prev = dll.Tail
	dll.Tail.Next = n
	dll.Tail = n
}

func (dll *DoublyLinkedList[T]) PopBack() (T, error) {
	if dll.Tail == nil {
		var z T
		return z, ErrorEmptyList
	}
	n := dll.Tail
	dll.Tail = n.Prev
	if dll.Tail == nil {
		dll.Head = nil
	} else {
		dll.Tail.Next = nil
	}
	return n.Data, nil
}

func (dll *DoublyLinkedList[T]) TopBack() (T, error) {
	if dll.Tail == nil {
		var z T
		return z, ErrorEmptyList
	}
	return dll.Tail.Data, nil
}

func (dll *DoublyLinkedList[T]) Empty() bool {
	return dll.Head == nil
}
