package main

import (
	"fmt"
	"github.com/hextechpal/dsa/lib/dset"
	"math"
	"sort"
)

type point struct {
	x, y int
}

func (p point) Key() point {
	return p
}

type edge struct {
	start point
	end   point
}

func (e edge) length() float64 {
	xdiff := e.end.x - e.start.x
	ydiff := e.end.y - e.start.y
	return math.Sqrt(float64(xdiff*xdiff + ydiff*ydiff))
}

type edges []edge

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i].length() < e[j].length()
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func connectingPoints(points []point) float64 {
	result := kruksalMST(points)
	total := float64(0)
	for i := 0; i < len(result); i++ {
		total += result[i].length()
	}
	return total
}

func kruksalMST(points []point) []edge {
	dSet := dset.NewDSet[point]()
	for _, point := range points {
		_ = dSet.MakeSet(point)
	}
	eges := make([]edge, 0)
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				eges = append(eges, edge{
					start: points[i],
					end:   points[j],
				})
			}
		}
	}

	sort.Sort(edges(eges))
	result := make([]edge, 0)
	for _, edge := range eges {
		if dSet.Find(edge.start) != dSet.Find(edge.end) {
			result = append(result, edge)
			dSet.Union(edge.start, edge.end)
		}
	}
	return result
}

func main() {
	points := []point{
		{0, 0},
		{0, 2},
		{1, 1},
		{3, 0},
		{3, 2},
	}
	fmt.Println(connectingPoints(points))
}
