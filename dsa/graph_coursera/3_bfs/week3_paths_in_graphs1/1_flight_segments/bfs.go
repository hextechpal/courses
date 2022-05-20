package main

import (
	"fmt"
	"github.com/hextechpal/dsa/lib/ll"
)

func shortestPathLength(graph [][]int, u, v int) int {
	dist := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		dist[i] = -1
	}
	deque := ll.SinglyLinkedList[int]{}
	deque.PushBack(u)
	dist[u] = 0

	for !deque.Empty() {
		node, _ := deque.PopFront()
		for _, adj := range graph[node] {
			if dist[adj] == -1 {
				deque.PushBack(adj)
				dist[adj] = dist[node] + 1
			}
		}
	}

	return dist[v]
}

func main() {
	graph := [][]int{
		{2, 3},
		{5},
		{0, 3},
		{0, 2},
		{1},
	}
	fmt.Println(shortestPathLength(graph, 3, 2))
}
