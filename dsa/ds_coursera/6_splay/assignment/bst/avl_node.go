package bst

type AvlNode struct {
	key    Comparer
	height int
	left   *AvlNode
	right  *AvlNode
	parent *AvlNode
}

func (n *AvlNode) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AvlNode) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

func (n *AvlNode) UpdateHeight() {
	if n.IsLeaf() {
		n.height = 1
	}
	maxH := n.left.Height()
	if maxH < n.right.Height() {
		maxH = n.right.Height()
	}
	n.height = 1 + maxH
}

func (n *AvlNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func NewNode(k Comparer) *AvlNode {
	return &AvlNode{
		key:    k,
		height: 1,
		left:   nil,
		right:  nil,
	}
}
