package heap

import (
	"fmt"
	"github.com/hextechpal/dsa/ds_coursera/lib/shared"
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	heap := NewMaxHeap[int](10, &shared.IntComparator{})
	_ = heap.Insert(5)
	_ = heap.Insert(4)
	_ = heap.Insert(10)
	_ = heap.Insert(1)
	_ = heap.Insert(7)
	_ = heap.Insert(14)
	_ = heap.Insert(12)

	fmt.Println(heap.GetMax())

	fmt.Println(heap.ExtractMax())

	fmt.Println(heap.GetMax())

	fmt.Println(heap.ExtractMax())

	fmt.Println(heap.GetMax())
}
