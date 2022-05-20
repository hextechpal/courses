package stack

import "errors"

var ErrorFullStack error = errors.New("stack is full")
var ErrorEmptyStack error = errors.New("stack is empty")

type Stack struct {
	arr  []any
	size int
	top  int
}

func NewStack(size int) *Stack {
	return &Stack{
		arr:  make([]any, size),
		size: size,
		top:  0,
	}
}

func (s *Stack) Push(e any) error {
	if s.Full() {
		return ErrorFullStack
	}
	s.arr[s.top] = e
	s.top++
	return nil
}

func (s *Stack) Pop() (any, error) {
	if s.Empty() {
		return nil, ErrorEmptyStack
	}
	el := s.arr[s.top-1]
	s.top--
	return el, nil
}

func (s *Stack) Peek() (any, error) {
	if s.Empty() {
		return nil, ErrorEmptyStack
	}
	return s.arr[s.top], nil
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Full() bool {
	return s.top == s.size
}

func (s *Stack) Empty() bool {
	return s.top == 0
}
