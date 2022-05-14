package main

import "fmt"

// binarySearch : This is a variation of binary search which returns the index of the element if found and a boolean
// to indicate the element is found
// If the element is not found it returns an index. The index indicated that if you need to insert this element in the
// array inorder to maintain this sorted this is the index where it should be inserted
func binarySearch(arr []int, val, start, end int) (int, bool) {
	if start <= end {
		mid := (start + end) / 2
		el := arr[mid]
		switch {
		case el == val:
			return mid, true
		case val < el:
			return binarySearch(arr, val, start, mid-1)
		default:
			return binarySearch(arr, val, mid+1, end)
		}
	} else {
		return start, false
	}
}

func main() {
	arr := []int{1, 5, 7, 9, 12, 14}
	idx, _ := binarySearch(arr, 20, 0, len(arr)-1)
	fmt.Println(idx)
}

