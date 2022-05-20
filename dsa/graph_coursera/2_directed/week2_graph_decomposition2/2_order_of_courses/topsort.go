package main

import "fmt"

type exploreMeta struct {
	visited bool
	pre     int
	post    int
}

func topSort(graph [][]int) []int {
	clock := 0
	order := make([]int, 0)
	meta := make(map[int]*exploreMeta)
	for node, _ := range graph {
		if m, ok := meta[node]; !ok || !m.visited {
			clock = explore(graph, node, meta, clock, &order)
		}
	}
	return order
}

func explore(graph [][]int, node int, meta map[int]*exploreMeta, clock int, order *[]int) int {
	meta[node] = &exploreMeta{
		visited: true,
		pre:     clock,
		post:    0,
	}
	clock += 1
	for _, child := range graph[node] {
		if m, ok := meta[child]; !ok || !m.visited {
			clock = explore(graph, child, meta, clock, order)
		}
	}
	meta[node].post = clock
	clock += 1
	*order = append(*order, node)
	return clock
}

func main() {
	graph := [][]int{
		{1},
		{2},
		{},
		{0},
	}
	fmt.Print(topSort(graph))
}
