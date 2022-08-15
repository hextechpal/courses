package main

import "fmt"

func binarySearch(arr []int, k int) int {
	return bs(arr, k, 0, len(arr)-1)
}

func bs(arr []int, k, l, r int) int {
	if l > r {
		return -1
	}
	m := l + (r-l)/2
	if arr[m] == k {
		return m
	} else if arr[m] > k {
		return bs(arr, k, l, m-1);
	} else {
		return bs(arr, k, m+1, r);
	}
}

func main() {
	arr := []int{1, 5, 8, 12, 13}

	fmt.Println(binarySearch(arr, 12))
	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 23))
}
