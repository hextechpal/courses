package max_stack

import (
	"errors"

	"github.com/hextechpal/dsa/ds_coursera/lib/stack"
)

var ErrorEmptyStack = errors.New("stack is empty")
var ErrorFullStack = errors.New("stack is full")

type MaxStack struct {
	data stack.Stack
	max  stack.Stack
}

func NewMaxStack (size int) *MaxStack{
	return &MaxStack{
		data: *stack.NewStack(size),
		max: *stack.NewStack(size),
	}
}

func (ms *MaxStack) Push(i int) error {
	err := ms.data.Push(i)
	if err != nil {
		return err
	}
	m ,_ := ms.max.Peek()
	if m.(int) < i{
		ms.max.Push(i)
	}else{
		ms.max.Push(m)		
	}
	return nil
}

func (ms *MaxStack) Pop() (int, error) {
	t, err := ms.data.Pop()
	if err != nil {
		return -1, nil
	}
	_, _ = ms.max.Pop()
	return t.(int), nil
}

func (ms *MaxStack) Max(i int) (int, error) {
	if ms.Empty() {
		return -1, ErrorEmptyStack
	}
	m, _ := ms.max.Peek()
	return m.(int), nil
}

func (ms *MaxStack) Peek() (int, error) {
	m, err := ms.max.Peek()
	if err != nil {
		return -1, err
	}
	return m.(int), nil
}

func (ms *MaxStack) Full() bool {
	return ms.data.Full()
}

func (ms *MaxStack) Empty() bool {
	return ms.data.Empty()
}
