package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/hextechpal/dsa/ds_coursera/lib/ll"
	"os"
	"strconv"
	"strings"
)

type idxValPair struct {
	idx int
	val int
}

func maxSliding(nums []int, wsize int) []int {
	result := make([]int, 0)
	deque := ll.DoublyLinkedList[idxValPair]{}
	for i, num := range nums {
		if deque.Empty() {
			deque.PushFront(idxValPair{
				idx: i,
				val: num,
			})
			continue
		}

		v, _ := deque.TopBack()
		if v.idx < i-wsize+1 {
			deque.PopBack()
		}

		for ivp, err := deque.TopFront(); ivp.val < num && err == nil; ivp, err = deque.TopFront() {
			deque.PopFront()
		}

		deque.PushFront(idxValPair{
			idx: i,
			val: num,
		})

		vf, _ := deque.TopBack()
		if i+1-wsize >= 0 {
			result = append(result, vf.val)
		}
	}

	return result
}

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

func main() {
	r := bufio.NewReader(os.Stdin)
	c, _ := r.ReadString('\n')
	count, _ := strconv.Atoi(strings.TrimSpace(c))
	in := make([]int, count)
	l, _ := r.ReadString('\n')

	nodes := strings.Split(strings.TrimSpace(l), " ")
	for i, n := range nodes {
		num, _ := strconv.Atoi(n)
		in[i] = num
	}

	l1, _ := r.ReadString('\n')
	w, _ := strconv.Atoi(strings.TrimSpace(l1))
	m := maxSliding(in, w)

	out := ""
	for _, max := range m {
		out += fmt.Sprintf("%d ", max)
	}

	fmt.Println(strings.TrimSpace(out))
}
