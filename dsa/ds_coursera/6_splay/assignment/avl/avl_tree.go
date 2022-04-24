package avl

import (
	"errors"
	"fmt"
)

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(k int, val interface{}) {
	t.root = t.root.add(k, val)
}

func (t *Tree) Rank(k int) int{
	return t.root.rank(k)
}

func (t *Tree) Find(k int) (interface{}, bool) {
	n := t.root.Find(k)
	if n == nil {
		return nil, false
	}
	return n.value, true
}

func (t *Tree) Delete(key int) error {
	n := t.root.delete(key)
	if n == nil {
		return errors.New("node not found")
	}
	return nil
}

func (t *Tree) InOrder() {
	t.root.inOrder()
}

type Node struct {
	key   int
	value interface{}

	height int
	size   int
	left   *Node
	right  *Node
}

func newNode(key int, val interface{}) *Node {
	return &Node{
		key:    key,
		value:  val,
		height: 1,
		size:   1,
		left:   nil,
		right:  nil,
	}
}

func (n *Node) UpdateHeight() {
	if n.IsLeaf() {
		n.height = 1
	}
	maxH := n.left.Height()
	if maxH < n.right.Height() {
		maxH = n.right.Height()
	}
	n.height = 1 + maxH
}

func (n *Node) UpdateSize() {
	n.size = n.left.Size() + n.right.Size() + 1
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (n *Node) add(key int, val interface{}) *Node {
	if n == nil {
		return newNode(key, val)
	}

	if n.key < key {
		n.right = n.right.add(key, val)
	} else if n.key > key {
		n.left = n.left.add(key, val)
	} else {
		n.value = val
	}
	return n.rebalance()
}

func (n *Node) rebalance() *Node {
	if n == nil {
		return nil
	}
	n.UpdateHeight()
	n.UpdateSize()
	bf := n.left.Height() - n.right.Height()
	if bf == 2 {
		l := n.left
		if l.right.Height() > l.left.Height() {
			n.left = l.rotateLeft()
		}
		return n.rotateRight()
	} else if bf == -2 {
		r := n.right
		if r.left.Height() > r.right.Height() {
			n.right = r.rotateRight()
		}
		return n.rotateLeft()
	} else {
		return n
	}
}

func (n *Node) rotateLeft() *Node {
	r := n.right
	n.right = r.left
	r.left = n
	n.UpdateHeight()
	n.UpdateSize()
	r.UpdateHeight()
	r.UpdateSize()
	return r
}

func (n *Node) rotateRight() *Node {
	l := n.left
	n.left = l.right
	l.right = n
	n.UpdateHeight()
	n.UpdateSize()
	l.UpdateHeight()
	l.UpdateSize()
	return l
}

func (n *Node) Find(k int) *Node {
	if n == nil {
		return nil
	}

	if n.key < k {
		return n.right.Find(k)
	} else if n.key > k {
		return n.left.Find(k)
	} else {
		return n
	}
}

func (n *Node) delete(key int) *Node {
	if n == nil {
		return nil
	}

	if n.key < key {
		n.right = n.right.delete(key)
	} else if n.key > key {
		n.left = n.left.delete(key)
	} else {
		if n.IsLeaf() {
			n = nil
			return nil
		} else if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.value = rightMinNode.value
			n.right = n.right.delete(rightMinNode.key)
		}
	}
	return n.rebalance()
}

func (n *Node) findSmallest() *Node {
	it := n
	for it.left != nil {
		it = it.left
	}
	return it
}

func (n *Node) inOrder() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.inOrder()
	}
	fmt.Printf("%d [%d %d] ", n.key, n.size, n.height)
	if n.right != nil {
		n.right.inOrder()
	}
}

func (n *Node) rank(k int) int {
	if n == nil{
		return -1
	}

	if n.key == k {
		return n.left.Size() + 1
	}else if n.key > k {
		return n.left.rank(k)
	}else{
		rr := n.right.rank(k)
		if rr == -1 {
			return rr
		}
		return (1 + n.left.Size()) + rr
	}
}
