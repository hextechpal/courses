package main

import (
	"fmt"
	"math"
)

func negativeCycle(graph [][]int, cost [][]int) bool {
	prev := make(map[int]int)
	dist := make(map[int]int)
	for v := range graph {
		dist[v] = math.MaxInt
	}
	dist[0] = 0
	previous := 0

	for i := 0; i < len(graph)-1; i++ {
		previous = -1
		for src, destinations := range graph {
			for _, dest := range destinations {
				if dist[dest] > dist[src]+cost[src][dest] && dist[src] != math.MaxInt {
					dist[dest] = dist[src] + cost[src][dest]
					prev[dest] = src
					previous = dest
				}
			}
		}
	}
	return previous != -1
}
func main() {
	graph := [][]int{
		{1},
		{2},
		{0},
		{0},
	}

	cost := [][]int{
		{0, -5, -1, -1},
		{-1, 0, 2, -1},
		{1, -1, 0, -1},
		{2, -1, -1, 0},
	}

	fmt.Println(negativeCycle(graph, cost))
}
