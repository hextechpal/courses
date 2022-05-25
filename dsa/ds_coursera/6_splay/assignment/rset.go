package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 1000000001

type avlNode struct {
	key    int64
	data   int64
	height int
	left   *avlNode
	right  *avlNode
	parent *avlNode
}

func (a *avlNode) insert(key int64, el int64, parent *avlNode) *avlNode {
	if a == nil {
		return &avlNode{
			key:    key,
			data:   el,
			height: 1,
			left:   nil,
			right:  nil,
			parent: parent,
		}
	}
	if a.key > key {
		a.left = a.left.insert(key, el, a)
	} else if a.key < key {
		a.right = a.right.insert(key, el, a)
	} else {
		return a
	}
	return a.rebalance()
}

func (a *avlNode) delete(key int64) *avlNode {
	if a == nil {
		return nil
	}
	if a.key > key {
		a.left = a.left.delete(key)
	} else if a.key < key {
		a.right = a.right.delete(key)
	} else {
		if a.isLeaf() {
			return nil
		} else if a.left == nil {
			return a.right
		} else if a.right == nil {
			return a.left
		} else {
			rightMinNode := a.right.leftDescendant()
			a.key = rightMinNode.key
			a.data = rightMinNode.data
			a.right = a.right.delete(rightMinNode.key)
		}
	}
	return a.rebalance()
}

func (a *avlNode) find(k int64) (*avlNode, bool) {
	if a == nil {
		return nil, false
	}
	if a.key == k {
		return a, true
	} else if a.key > k {
		if a.left == nil {
			return a, false
		}
		return a.left.find(k)
	} else {
		if a.right == nil {
			return a, false
		}
		return a.right.find(k)
	}
}

func (a *avlNode) rebalance() *avlNode {
	a.updateHeight()
	bf := a.left.getHeight() - a.right.getHeight()
	if bf == -2 {
		// This means right subtree is heavier
		return a.rebalanceRight()
	} else if bf == 2 {
		// This means left subtree is heavier
		return a.rebalanceLeft()
	} else {
		return a
	}
}

func (a *avlNode) rebalanceLeft() *avlNode {
	l := a.left
	if l.left.getHeight() < l.right.getHeight() {
		// LR imbalance
		a.left = l.rotateLeft()
	}
	// LL imbalance
	return a.rotateRight()

}

func (a *avlNode) rebalanceRight() *avlNode {
	r := a.right
	if r.right.getHeight() < r.left.getHeight() {
		// RL imbalance
		a.right = r.rotateRight()
	}
	// RR imbalance
	return a.rotateLeft()
}

func (a *avlNode) rotateLeft() *avlNode {
	r := a.right

	a.right = r.left
	if r.left != nil {
		r.left.parent = a
	}

	r.left = a
	r.parent = a.parent
	a.parent = r

	a.updateHeight()
	r.updateHeight()
	return r

}

func (a *avlNode) rotateRight() *avlNode {
	l := a.left
	a.left = l.right
	if l.right != nil {
		l.right.parent = a
	}

	l.right = a
	l.parent = a.parent
	a.parent = l

	a.updateHeight()
	l.updateHeight()
	return l
}

func (a *avlNode) getHeight() int {
	if a == nil {
		return 0
	}
	return a.height
}

func (a *avlNode) updateHeight() {
	if a == nil {
		return
	}
	maxHeight := a.left.getHeight()
	if maxHeight < a.right.getHeight() {
		maxHeight = a.right.getHeight()
	}
	a.height = 1 + maxHeight

}

func (a *avlNode) isLeaf() bool {
	if a == nil {
		return true
	}
	return a.left == nil && a.right == nil
}

func (a *avlNode) leftDescendant() *avlNode {
	if a.left == nil {
		return a
	}
	return a.left.leftDescendant()
}

func (a *avlNode) rightAncestor() *avlNode {
	if a == nil || a.parent == nil {
		return nil
	}
	if a.key < a.parent.key {
		return a.parent
	}
	return a.parent.rightAncestor()
}

func (a *avlNode) next() *avlNode {
	if a == nil {
		return nil
	}

	if a.right != nil {
		return a.right.leftDescendant()
	} else {
		return a.rightAncestor()
	}
}

type Strategy string

const (
	INORDER   Strategy = "inorder"
	PREORDER  Strategy = "preorder"
	POSTORDER Strategy = "postorder"
)

type AVLTree struct {
	root *avlNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(k int64, el int64) {
	t.root = t.root.insert(k, el, t.root)
}

func (t *AVLTree) Find(key int64) (int64, error) {
	n, ok := t.root.find(key)
	if !ok {
		return -1, errors.New("not found")
	}
	return n.data, nil
}

func (t *AVLTree) RangeSearch(low, high int64) []int64 {
	result := make([]int64, 0)
	n, _ := t.root.find(low)
	for n != nil && n.key <= high {
		if n.key >= low {
			result = append(result, n.data)
		}
		n = n.next()
	}
	return result
}

func (t AVLTree) Traverse(strategy Strategy) []int64 {
	result := make([]int64, 0)
	switch strategy {
	case INORDER:
		traverseInOrder(t.root, &result)
	case POSTORDER:
		traversePostOrder(t.root, &result)
	case PREORDER:
		traversePreOrder(t.root, &result)
	}
	return result
}

func (t *AVLTree) Delete(key int64) {
	t.root = t.root.delete(key)
}

func traversePostOrder(root *avlNode, result *[]int64) {
	if root.left != nil {
		traversePostOrder(root.left, result)
	}

	if root.right != nil {
		traversePostOrder(root.right, result)
	}
	*result = append(*result, root.data)
}

func traversePreOrder(root *avlNode, result *[]int64) {
	*result = append(*result, root.data)

	if root.left != nil {
		traversePreOrder(root.left, result)
	}

	if root.right != nil {
		traversePreOrder(root.right, result)
	}
}

func traverseInOrder(root *avlNode, result *[]int64) {
	if root.left != nil {
		traverseInOrder(root.left, result)
	}
	*result = append(*result, root.data)
	if root.right != nil {
		traverseInOrder(root.right, result)
	}
}

type RSet struct {
	t *AVLTree
	x int64
}

func NewRSet() *RSet {
	return &RSet{
		t: NewAVLTree(),
		x: 0,
	}
}

func (r *RSet) Add(i int64) {
	val := (r.x + i) % M
	r.t.Insert(val, val)
}

func (r *RSet) Del(i int64) {
	val := (r.x + i) % M
	r.t.Delete(val)
}

func (r *RSet) Find(i int64) {
	val := (r.x + i) % M
	_, err := r.t.Find(val)
	if err != nil {
		fmt.Println("Not found")
		return
	}
	fmt.Println("Found")

}

func (r *RSet) Sum(left, right int64) {
	l := (r.x + left) % M
	h := (r.x + right) % M

	elements := r.t.RangeSearch(l, h)
	sum := int64(0)
	for _, el := range elements {
		sum += el
	}
	r.x = sum
	fmt.Println(sum)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l1, _ := r.ReadString('\n')

	ops, _ := strconv.Atoi(strings.Trim(l1, "\n"))

	cmds := make([]string, ops)
	for i := 0; i < ops; i++ {
		op, _ := r.ReadString('\n')
		cmds[i] = op
	}
	rset := NewRSet()
	for _, cmd := range cmds {
		args := strings.Split(strings.Trim(cmd, "\n"), " ")
		switch args[0] {
		case "+":
			add, _ := strconv.ParseInt(args[1], 10, 64)
			rset.Add(add)
		case "-":
			del, _ := strconv.ParseInt(args[1], 10, 64)
			rset.Del(del)
		case "?":
			find, _ := strconv.ParseInt(args[1], 10, 64)
			rset.Find(find)
		case "s":
			left, _ := strconv.ParseInt(args[1], 10, 64)
			right, _ := strconv.ParseInt(args[2], 10, 64)
			rset.Sum(left, right)
		}
	}
}
