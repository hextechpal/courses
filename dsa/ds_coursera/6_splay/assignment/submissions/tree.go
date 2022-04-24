package submissions

import (
	"fmt"
	"math"
)

type Tree struct {
	keys   []int
	lefts  []int
	rights []int
}

func (t *Tree) InOrder() {
	t.inOrder(0)
}

func (t *Tree) inOrder(i int) {
	if i < len(t.lefts) && t.lefts[i] != -1 {
		t.inOrder(t.lefts[i])
	}
	fmt.Printf("%d ", t.keys[i])
	if i < len(t.rights) && t.rights[i] != -1 {
		t.inOrder(t.rights[i])
	}
}

func (t *Tree) PreOrder() {
	t.preOrder(0)
}

func (t *Tree) preOrder(i int) {
	fmt.Printf("%d ", t.keys[i])
	if i < len(t.lefts) && t.lefts[i] != -1 {
		t.preOrder(t.lefts[i])
	}
	if i < len(t.rights) && t.rights[i] != -1 {
		t.preOrder(t.rights[i])
	}
}

func (t *Tree) PostOrder() {
	t.postOrder(0)
}

func (t *Tree) postOrder(i int) {
	if i < len(t.lefts) && t.lefts[i] != -1 {
		t.postOrder(t.lefts[i])
	}
	if i < len(t.rights) && t.rights[i] != -1 {
		t.postOrder(t.rights[i])
	}
	fmt.Printf("%d ", t.keys[i])
}

func (t *Tree) IsBST() bool {
	return t.isBST(0, math.MinInt, math.MaxInt)
}

func (t *Tree) isBST(i int, min, max int) bool {
	if i == -1 {
		return true
	}
	if t.keys[i] <= min || t.keys[i] > max {
		return false
	}

	return t.isBST(t.lefts[i], min, t.keys[i]) && t.isBST(t.rights[i], t.keys[i], max)
}

func NewTree(keys, lefts, rights []int) *Tree {
	return &Tree{keys: keys, lefts: lefts, rights: rights}
}
