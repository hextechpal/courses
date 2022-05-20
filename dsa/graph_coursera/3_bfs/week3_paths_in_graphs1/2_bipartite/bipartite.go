package main

import (
	"fmt"
	"github.com/hextechpal/dsa/lib/ll"
)

func bipartite(graph [][]int) bool {
	dist := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		dist[i] = -1
	}
	deque := ll.SinglyLinkedList[int]{}
	deque.PushBack(0)
	dist[0] = 0

	for !deque.Empty() {
		node, _ := deque.PopFront()
		for _, adj := range graph[node] {
			if dist[adj] == -1 {
				deque.PushBack(adj)
				dist[adj] = dist[node] + 1
			} else if dist[adj] == dist[node] {
				return false
			}
		}
	}

	for _, v := range dist {
		if v == -1 {
			return false
		}
	}

	return true
}

func main() {
	graph := [][]int{
		{3},
		{3, 4},
		{3, 4},
		{0, 1, 2},
		{1},
	}
	fmt.Println(bipartite(graph))
}
