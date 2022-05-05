package ll

// SinglyLinkedList with Tail
type SinglyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// PushFront Pushes to the front of the LinkedList.
// O(1)
func (sll *SinglyLinkedList[T]) PushFront(data T) {
	node := &Node[T]{Data: data}
	node.Next = sll.head
	sll.head = node
	if sll.tail == nil {
		sll.tail = sll.head
	}
}

// TopFront Provides Data from the Head
// O(1)
func (sll *SinglyLinkedList[T]) TopFront() (T, error) {
	if sll.head == nil {
		var r T
		return r, ErrorEmptyList
	}
	return sll.head.Data, nil
}

func (sll *SinglyLinkedList[T]) PopFront() (T, error) {
	if sll.head == nil {
		var z T
		return z, ErrorEmptyList
	}
	n := sll.head
	sll.head = sll.head.Next
	if sll.head == nil {
		sll.tail = nil
	}
	return n.Data, nil
}

// PushFront Pushes to the front of the LinkedList.
// O(1)
func (sll *SinglyLinkedList[T]) PushBack(data T) {
	node := &Node[T]{Data: data}
	if sll.tail == nil {
		sll.tail = node
		sll.head = node
		return
	}
	sll.tail.Next = node
	sll.tail = node
}

func (sll *SinglyLinkedList[T]) TopBack() (T, error) {
	if sll.tail == nil {
		var z T
		return z, ErrorEmptyList
	}

	return sll.tail.Data, nil
}

func (sll *SinglyLinkedList[T]) PopBack() (T, error) {
	if sll.tail == nil {
		var z T
		return z, ErrorEmptyList
	}

	if sll.head == sll.tail {
		n := sll.head
		sll.head = nil
		sll.tail = nil
		return n.Data, nil
	}

	curr := sll.head
	prev := sll.head

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}
	prev.Next = nil
	sll.tail = prev
	return curr.Data, nil
}

func (sll *SinglyLinkedList[T]) Empty() bool {
	return sll.head == nil
}
