package main

import "fmt"

func scc(graph [][]int) int {
	post := calculatePostOrder(graph)

	visited := make(map[int]bool)
	sccList := make([][]int, 0)
	for i := len(post) - 1; i >= 0; i-- {
		if _, ok := visited[post[i]]; !ok {
			scc := explore(graph, post[i], visited, make([]int, 0))
			sccList = append(sccList, scc)
		}
	}
	return len(sccList)
}

func calculatePostOrder(graph [][]int) []int {
	rgraph := reverseGraph(graph)
	visited := make(map[int]bool)
	post := make([]int, 0)

	for i := range graph {
		if _, ok := visited[i]; !ok {
			post = explore(rgraph, i, visited, post)
		}
	}
	return post
}

func explore(graph [][]int, node int, visited map[int]bool, post []int) []int {
	visited[node] = true

	for _, child := range graph[node] {
		if _, ok := visited[child]; !ok {
			post = explore(graph, child, visited, post)
		}
	}
	post = append(post, node)
	return post
}

func reverseGraph(graph [][]int) [][]int {
	rgraph := make([][]int, len(graph))
	for i := 0; i < len(rgraph); i++ {
		rgraph[i] = make([]int, 0)
	}

	for i, children := range graph {
		for _, child := range children {
			rgraph[child] = append(rgraph[child], i)
		}
	}
	return rgraph
}

func main() {
	graph := [][]int{
		{},
		{0},
		{0, 1},
		{0, 2},
		{1, 2},
	}
	fmt.Println(scc(graph))
}
