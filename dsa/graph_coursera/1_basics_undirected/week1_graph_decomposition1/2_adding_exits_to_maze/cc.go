package main

import "fmt"

func connectedComponents(graph [][]int) int {
	cc := 0
	visited := make(map[int]bool)
	for i, _ := range graph {
		if _, ok := visited[i]; !ok {
			explore(graph, i, visited)
			cc += 1
		}
	}
	return cc
}

func explore(graph [][]int, a int, visited map[int]bool) {
	visited[a] = true
	for _, node := range graph[a] {
		if _, ok := visited[node]; !ok {
			explore(graph, node, visited)
		}
	}
}

func main() {
	graph := [][]int{
		{1},
		{0, 2},
		{1},
		{},
	}
	fmt.Print(connectedComponents(graph))
}
