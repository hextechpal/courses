package main

import "fmt"

type counter struct {
	arr   []int
	count int
}

func (c *counter) countInversions() {
	c.ciInternal(0, len(c.arr)-1)
}

func (c *counter) ciInternal(l, r int) {
	if l >= r {
		return
	}

	m := l + (r-l)/2

	c.ciInternal(l, m)
	c.ciInternal(m+1, r)

	c.merge(l, m, r)
}

func (c *counter) merge(l, m, r int) {
	lsize := m - l + 1
	rsize := r - m

	larr := make([]int, lsize)
	rarr := make([]int, rsize)

	for i := 0; i < lsize; i++ {
		larr[i] = c.arr[l+i]
	}

	for i := 0; i < rsize; i++ {
		rarr[i] = c.arr[m+i+1]
	}

	i := 0
	j := 0

	k := l

	for i < lsize && j < rsize {
		if larr[i] < rarr[j] {
			c.arr[k] = larr[i]
			i++
		} else {
			c.arr[k] = rarr[j]
			j++
			c.count++
		}
		k++
	}

	for i < lsize {
		c.arr[k] = larr[i]
		k++
		i++
	}

	for j < rsize {
		c.arr[k] = rarr[j]
		k++
		j++
	}
}

func countInversions(arr []int) int {
	c := &counter{
		arr:   arr,
		count: 0,
	}
	c.countInversions()
	return c.count
}

func main() {
	fmt.Println(countInversions([]int{5, 2, 6, 1}))
}
