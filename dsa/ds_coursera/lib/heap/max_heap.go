package heap

import (
	"errors"
	"github.com/hextechpal/dsa/ds_coursera/lib/shared"
)

type MaxHeap[T any] struct {
	capacity   int
	size       int
	sl         []T
	comparator shared.Comparator[T]
}

func NewMaxHeap[T any](capacity int, comparator shared.Comparator[T]) *MaxHeap[T] {
	return &MaxHeap[T]{size: 0, capacity: capacity, sl: make([]T, capacity), comparator: comparator}
}

func (mh *MaxHeap[T]) GetMax() (T, error) {
	if mh.size == 0 {
		var z T
		return z, errors.New("empty heap")
	}
	return mh.sl[0], nil
}

func (mh *MaxHeap[T]) ExtractMax() (T, error) {
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

func (mh *MaxHeap[T]) Insert(el T) error {
	if mh.size == mh.capacity {
		return errors.New("heap full")
	}
	mh.sl[mh.size] = el
	mh.size++
	mh.shiftUp(mh.size - 1)
	return nil
}

func (mh *MaxHeap[_]) parent(i int) int {
	return (i - 1) / 2
}

func (mh *MaxHeap[_]) leftChild(i int) int {
	return 2*i + 1
}

func (mh *MaxHeap[_]) rightChild(i int) int {
	return 2*i + 2
}

func (mh *MaxHeap[T]) shiftUp(i int) {
	if i == 0 {
		return
	}
	parent := mh.parent(i)
	if mh.comparator.Compare(mh.sl[parent], mh.sl[i]) < 0 {
		mh.sl[parent], mh.sl[i] = mh.sl[i], mh.sl[parent]
		mh.shiftUp(parent)
	}
}

func (mh *MaxHeap[T]) shiftDown(i int) {
	maxIdx := i
	lc := mh.leftChild(i)

	if lc < mh.size && mh.comparator.Compare(mh.sl[maxIdx], mh.sl[lc]) < 0 {
		maxIdx = lc
	}

	rc := mh.rightChild(i)
	if rc < mh.size && mh.comparator.Compare(mh.sl[maxIdx], mh.sl[rc]) < 0 {
		maxIdx = rc
	}

	if maxIdx != i {
		mh.sl[maxIdx], mh.sl[i] = mh.sl[i], mh.sl[maxIdx]
		mh.shiftDown(maxIdx)
	}

}
