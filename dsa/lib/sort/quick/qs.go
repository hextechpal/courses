package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) {
	qsInternal(&arr, 0, len(arr)-1)
}

func qsInternal(arr *[]int, l int, r int) {
	if l >= r {
		return
	}
	k := l + rand.Intn(r-l)
	(*arr)[k], (*arr)[l] = (*arr)[l], (*arr)[k]
	m := partition2(arr, l, r)
	qsInternal(arr, l, m-1)
	qsInternal(arr, m+1, r)
}

func partition2(arr *[]int, l int, r int) int {
	pivot := (*arr)[l]
	j := l

	for i := l + 1; i <= r; i++ {
		if (*arr)[i] <= pivot {
			j++
			(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
		}
	}

	(*arr)[l], (*arr)[j] = (*arr)[j], (*arr)[l]
	return j
}

func main() {
	arr := []int{6, 6, 7, 6, 4, 1, 8, 6, 1}
	quickSort(arr)
	fmt.Println(arr)
}
