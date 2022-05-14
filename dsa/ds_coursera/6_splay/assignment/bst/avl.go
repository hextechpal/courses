package bst

import (
	"errors"
	"fmt"
)

type AvlTree struct {
	root *AvlNode
}

func (t *AvlTree) Search(k Comparer) Comparer {
	node := find(k, t.root)
	if node == nil {
		return nil
	}
	return node.key
}

func (t *AvlTree) Find(k Comparer) *AvlNode {
	return find(k, t.root)
}

func find(k Comparer, node *AvlNode) *AvlNode {
	if node == nil || k == nil {
		return nil
	}
	comp := node.key.Compare(k)
	if comp == 0 {
		return node
	}

	if comp == 1 {
		return find(k, node.left)
	} else {
		return find(k, node.right)
	}
}

func (t *AvlTree) InOrder() []int {
	order := inOrder(t.root, make([]int, 0))
	fmt.Println()
	return order
}

func inOrder(node *AvlNode, ints []int) []int {
	if node.left != nil {
		ints = inOrder(node.left, ints)
	}
	fmt.Printf("%d(%d) ", node.key.GetValue(), node.height)
	ints = append(ints, node.key.GetValue())
	if node.right != nil {
		ints = inOrder(node.right, ints)
	}
	return ints
}

func (t *AvlTree) AvlInsert(k Comparer) {
	if t.root == nil {
		t.root = NewNode(k)
		return
	}
	insert(k, t.root)
	n := t.Find(k)
	rebalance(n, t)
}

func (t *AvlTree) AvlDelete(k Comparer) error {
	return deleteNode(k, t.root)
}

func deleteNode(k Comparer, n *AvlNode) error {
	if n == nil {
		return errors.New("node not present")
	}
	cmp := n.key.Compare(k)
	if cmp == 1 {
		return deleteNode(k, n.left)
	} else if cmp == -1 {
		return deleteNode(k, n.left)
	} else {
		if n.IsLeaf() {
			if n.parent.left == n {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
			n = nil
		} else if n.left == nil {
			r := n.right
			if n.parent.left == n {
				n.parent.left = r
			} else {
				n.parent.right = r
			}
			n = nil
		} else if n.right == nil {
			l := n.left
			if n.parent.left == n {
				n.parent.left = l
			} else {
				n.parent.right = l
			}
			n = nil
		} else {
			p := n.parent
			if p.left == n {
				p.left = n.right
				n.right.parent = p
			} else {
				p.right = n.left
				n.left.parent = p
			}
			n = nil
		}
		return nil
	}
}

func insert(k Comparer, node *AvlNode) {
	cmp := node.key.Compare(k)
	if cmp > 0 {
		if node.left != nil {
			insert(k, node.left)
		} else {
			node.left = NewNode(k)
		}
	} else {
		if node.right != nil {
			insert(k, node.right)
		} else {
			node.right = NewNode(k)
		}
	}
	node.UpdateHeight()
}

func rebalance(n *AvlNode, t *AvlTree) {
	p := n.parent
	bf := n.BalanceFactor()
	if bf > 1 {
		// This means left side is heavier
		rebalanceRight(n, t)
	} else if bf < -1 {
		// This means right side is heavier
		rebalanceLeft(n, t)
	}
	if p != nil {
		rebalance(p, t)
	}
}

func rebalanceLeft(n *AvlNode, t *AvlTree) {
	r := n.right
	if r.left.Height() > r.right.Height() {
		rotateRight(r, t)
	}
	rotateLeft(n, t)
}

func rebalanceRight(n *AvlNode, t *AvlTree) {
	l := n.left
	if l.right.Height() > l.left.Height() {
		rotateLeft(l, t)
	}
	rotateRight(n, t)
}

func rotateRight(n *AvlNode, t *AvlTree) {
	l := n.left
	n.left = l.right
	if l.right != nil {
		l.right.parent = n
	}
	l.right = n
	l.parent = n.parent
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = l
		} else {
			n.parent.right = l
		}
	} else {
		t.root = l
	}
	n.parent = l
	n.UpdateHeight()
	l.UpdateHeight()

}

func rotateLeft(n *AvlNode, t *AvlTree) {
	r := n.right
	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}
	r.left = n
	r.parent = n.parent
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = r
		} else {
			n.parent.right = r
		}
	} else {
		t.root = r
	}
	n.parent = r
	n.UpdateHeight()
	r.UpdateHeight()
}
