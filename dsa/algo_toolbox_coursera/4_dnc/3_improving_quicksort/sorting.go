package main

import (
	"fmt"
	"math/rand"
)

func qs3(arr []int) {
	qs3Internal(&arr, 0, len(arr)-1)
}

func qs3Internal(arr *[]int, l, r int) {
	if l >= r {
		return
	}
	k := l + rand.Intn(r-l)
	(*arr)[k], (*arr)[l] = (*arr)[l], (*arr)[k]

	m1, m2 := partition3(arr, l, r)
	qs3Internal(arr, l, m1-1)
	qs3Internal(arr, m2+1, r)
}

func partition3(arr *[]int, l int, r int) (int, int) {
	pivot := (*arr)[l]
	m1 := l
	m2 := r
	i := l + 1

	for i <= m2 {
		if (*arr)[i] < pivot {
			m1++
			(*arr)[m1], (*arr)[i] = (*arr)[i], (*arr)[m1]
			i++
		} else if (*arr)[i] > pivot {
			(*arr)[m2], (*arr)[i] = (*arr)[i], (*arr)[m2]
			m2--
		} else {
			i++
		}
	}
	(*arr)[l], (*arr)[m1] = (*arr)[m1], (*arr)[l]
	return m1, m2
}

func main() {
	arr := []int{6, 6, 7, 6, 4, 1, 8, 6, 1}
	qs3(arr)

	fmt.Println(arr)
}
