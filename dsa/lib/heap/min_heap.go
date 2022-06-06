package heap

import (
	"errors"
	"github.com/hextechpal/dsa/lib/shared"
)

type MinHeap[T any] struct {
	capacity   int
	size       int
	sl         []T
	comparator shared.Comparator[T]
}

func NewMinHeap[T any](capacity int, comparator shared.Comparator[T]) *MinHeap[T] {
	return &MinHeap[T]{size: 0, capacity: capacity, sl: make([]T, capacity), comparator: comparator}
}

func (mh *MinHeap[T]) GetMin() (T, error) {
	if mh.size == 0 {
		var z T
		return z, errors.New("empty heap")
	}
	return mh.sl[0], nil
}

func (mh *MinHeap[T]) ChangePriority(i int, el T) {
	oldEl := mh.sl[i]
	mh.sl[i] = el
	if mh.comparator.Compare(oldEl, el) < 0 {
		mh.shiftDown(i)
	} else {
		mh.shiftUp(i)
	}
}

func (mh *MinHeap[T]) ExtractMin() (T, error) {
	if mh.size == 0 {
		var z T
		return z, errors.New("empty heap")
	}
	root := mh.sl[0]
	mh.sl[0] = mh.sl[mh.size-1]
	mh.shiftDown(0)
	mh.size--
	return root, nil
}

func (mh *MinHeap[T]) Insert(el T) error {
	if mh.size == mh.capacity {
		return errors.New("heap full")
	}
	mh.sl[mh.size] = el
	mh.size++
	mh.shiftUp(mh.size - 1)
	return nil
}

func (mh *MinHeap[_]) Empty() bool {
	return mh.size == 0
}

func (mh *MinHeap[_]) parent(i int) int {
	return (i - 1) / 2
}

func (mh *MinHeap[_]) leftChild(i int) int {
	return 2*i + 1
}

func (mh *MinHeap[_]) rightChild(i int) int {
	return 2*i + 2
}

func (mh *MinHeap[T]) shiftUp(i int) {
	if i == 0 {
		return
	}
	parent := mh.parent(i)
	if mh.comparator.Compare(mh.sl[parent], mh.sl[i]) > 0 {
		mh.sl[parent], mh.sl[i] = mh.sl[i], mh.sl[parent]
		mh.shiftUp(parent)
	}
}

func (mh *MinHeap[T]) shiftDown(i int) {
	minIdx := i
	lc := mh.leftChild(i)

	if lc < mh.size && mh.comparator.Compare(mh.sl[minIdx], mh.sl[lc]) > 0 {
		minIdx = lc
	}

	rc := mh.rightChild(i)
	if rc < mh.size && mh.comparator.Compare(mh.sl[minIdx], mh.sl[rc]) > 0 {
		minIdx = rc
	}
	if minIdx != i {
		mh.sl[minIdx], mh.sl[i] = mh.sl[i], mh.sl[minIdx]
		mh.shiftDown(minIdx)
	}
}
