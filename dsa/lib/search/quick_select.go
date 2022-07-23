package main

import "errors"

// QuickSelect : Quick Select selects the kth smallest number from the array and returns it
func QuickSelect(arr []int, k int) (int, error) {
	actualK := k
	if k < 0 {
		actualK = len(arr) + 1 + k
	}

	if len(arr) < actualK {
		return -1, errors.New("k cannot be larger than array length")
	}

	return qSelect(arr, actualK, 0, len(arr)-1), nil
}

func qSelect(arr []int, k int, left int, right int) int {
	if left == right {
		return arr[left]
	}

	pivot := partition(arr, left, right)

	if pivot == k-1 {
		return arr[pivot]
	} else if pivot < k {
		return qSelect(arr, k, pivot+1, right)
	} else {
		return qSelect(arr, k, left, pivot-1)
	}
}

func partition(arr []int, left int, right int) int {
	pivotVal := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] < pivotVal {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
