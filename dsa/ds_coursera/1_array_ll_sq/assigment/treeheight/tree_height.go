package treeheight


type Tree struct{
	n int
	parents []int
}

func (t *Tree) Height() int{
	height := 0
	depths := make([]int, t.n)
	for i:= 0; i<t.n; i++{
		depths[i] = t.depth(i, depths)
		if height< depths[i] {
			height = depths[i]
		}
	}
	return height
}

func (t *Tree) depth(node int, depths []int) int{
	if depths[node] != 0{
		return depths[node]
	}

	if t.parents[node] == -1{
		return 1
	}

	return 1 + t.depth(t.parents[node], depths)
}