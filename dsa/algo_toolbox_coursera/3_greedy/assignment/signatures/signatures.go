package main

import "fmt"

type segment struct {
	x, y int
}

func minPoints(segments []segment) []int {
	if len(segments) == 0 {
		return []int{}
	}

	r := make([]int, 0)
	i := len(segments) - 1
	r = append(r, segments[i].x)

	for i >= 0 {
		for i >= 0 && segments[i].y >= r[len(r)-1] {
			i--
		}

		if i >= 0 {
			r = append(r, segments[i].x)
		}
	}

	return r
}
func main() {
	//p := minPoints([]segment{{1, 3}, {2, 5}, {3, 6}})
	p := minPoints([]segment{{1, 3}, {2, 4}, {4, 7}, {5, 6}})
	//p := minPoints([]segment{{1, 3}})
	fmt.Println(p)
}
