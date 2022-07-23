package main

// BinarySearch : This is a variation of binary search which returns the index of the element if found and a boolean
// to indicate the element is found
// If the element is not found it returns an index. The index indicated that if you need to insert this element in the
// array inorder to maintain this sorted this is the index where it should be inserted
func BinarySearch(arr []int, val, start, end int) (int, bool) {
	if start <= end {
		mid := (start + end) / 2
		el := arr[mid]
		switch {
		case el == val:
			return mid, true
		case val < el:
			return BinarySearch(arr, val, start, mid-1)
		default:
			return BinarySearch(arr, val, mid+1, end)
		}
	} else {
		return start, false
	}
}
