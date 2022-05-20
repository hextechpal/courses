package main

import "fmt"

func isCyclic(graph [][]int) bool {
	visited := make(map[int]bool)
	stack := make(map[int]bool)

	for node, _ := range graph {
		if _, ok := visited[node]; !ok {
			cyclic := explore(graph, node, stack, visited)
			if cyclic {
				return true
			}
		}
	}
	return false
}

func explore(graph [][]int, node int, stack map[int]bool, visited map[int]bool) bool {
	if stack[node] {
		return true
	}

	if visited[node] {
		return false
	}
	stack[node] = true
	visited[node] = true

	for _, child := range graph[node] {
		if explore(graph, child, stack, visited) {
			return true
		}
	}
	stack[node] = false
	return false
}

func main() {
	graph := [][]int{
		{1},
		{2},
		{0},
		{0},
	}
	fmt.Print(isCyclic(graph))
}
