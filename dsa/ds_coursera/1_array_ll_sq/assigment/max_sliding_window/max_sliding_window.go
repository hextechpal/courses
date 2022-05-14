package main

import (
	"github.com/hextechpal/dsa/ds_coursera/lib/ll"
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
