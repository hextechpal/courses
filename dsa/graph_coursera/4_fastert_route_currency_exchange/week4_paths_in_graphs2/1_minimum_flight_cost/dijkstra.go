package main

import (
	"fmt"
	"math"
)

func shortestPath(graph [][]int, cost [][]int, x, y int) int {
	m, p := dijkstra(graph, cost, x)
	fmt.Println(p)
	return m[y]
}

func dijkstra(graph [][]int, cost [][]int, x int) (map[int]int, map[int]int) {
	dist := make(map[int]int)
	visited := make(map[int]bool)
	prev := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		dist[i] = -1
		visited[i] = false
	}
	dist[x] = 0
	prev[x] = x

	for !allVisited(visited) {
		minVertex := extractMin(dist, visited)
		visited[minVertex] = true
		for _, vertex := range graph[minVertex] {
			if dist[vertex] == -1 || dist[vertex] > dist[minVertex]+cost[minVertex][vertex] {
				dist[vertex] = dist[minVertex] + cost[minVertex][vertex]
				prev[vertex] = minVertex
			}
		}
	}

	return dist, prev
}

func extractMin(dist map[int]int, visited map[int]bool) int {
	minDist := math.MaxInt
	minVertex := -1
	for vertex, dist := range dist {
		if minVertex == -1 {
			minVertex = vertex
		}
		if dist != -1 && minDist > dist && !visited[vertex] {
			minDist = dist
			minVertex = vertex
		}
	}
	return minVertex
}

func allVisited(visited map[int]bool) bool {
	if len(visited) == 0 {
		return false
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	graph := [][]int{
		{1, 2},
		{},
		{1},
	}

	cost := [][]int{
		{0, 5, 6},
		{-1, 0, -1},
		{-1, -3, 0},
	}
	fmt.Println(shortestPath(graph, cost, 0, 2))
}
