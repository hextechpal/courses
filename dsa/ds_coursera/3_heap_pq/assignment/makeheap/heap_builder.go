package makeheap

import (
	"errors"
	"fmt"
)

type swap struct {
	to   int
	from int
}

type HeapBuilder struct {
	maxSize int
	size    int
	data    []int
	swaps   []swap
}

func (h *HeapBuilder) Insert(el int) error {
	if h.size >= h.maxSize {
		return errors.New("max heap size reached")
	}
	h.size++
	h.data[h.size-1] = el
	h.shiftUp(h.size - 1)
	return nil
}

func (h *HeapBuilder) Heapify() {
	lastNonLeaf := h.parent(h.size - 1)
	for i := lastNonLeaf; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *HeapBuilder) shiftUp(i int) {
	for i > 0 && h.data[h.parent(i)] > h.data[i] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *HeapBuilder) shiftDown(i int) {
	minIndex := i
	lc := h.leftChild(i)
	if lc < h.maxSize && h.data[lc] < h.data[minIndex] {
		minIndex = lc
	}
	rc := h.rightChild(i)
	if rc < h.maxSize && h.data[rc] < h.data[minIndex] {
		minIndex = rc
	}
	if i != minIndex {
		h.swap(i, minIndex)
		h.shiftDown(minIndex)
	}
}

func (h *HeapBuilder) swap(from int, to int) {
	h.swaps = append(h.swaps, swap{to: to, from: from})
	h.data[from], h.data[to] = h.data[to], h.data[from]
}

func (h *HeapBuilder) parent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}

func (h *HeapBuilder) rightChild(i int) int {
	return 2*i + 2
}

func (h *HeapBuilder) leftChild(i int) int {
	return 2*i + 1
}

func (h *HeapBuilder) PrintSwaps() {
	fmt.Println(len(h.swaps))
	for _, sw := range h.swaps {
		fmt.Printf("%d %d\n", sw.from, sw.to)
	}
}

func NewBuildHeap(data []int) *HeapBuilder {
	return &HeapBuilder{
		maxSize: cap(data),
		size:    len(data),
		data:    data,
		swaps:   make([]swap, 0),
	}
}
