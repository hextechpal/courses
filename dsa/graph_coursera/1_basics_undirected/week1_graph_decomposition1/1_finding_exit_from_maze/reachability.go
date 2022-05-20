package main

import "fmt"

func reach(graph [][]int, a, b int) bool {
	visited := make(map[int]bool)
	explore(graph, a, visited)
	return visited[b]
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
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}
	if reach(graph, 0, 3) {
		fmt.Print(1)
		return
	}
	fmt.Print(0)
}
