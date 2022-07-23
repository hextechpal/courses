package tree

import "errors"

type avlNode[N any] struct {
	key    int
	data   N
	height int
	left   *avlNode[N]
	right  *avlNode[N]
	parent *avlNode[N]
}

func (a *avlNode[N]) insert(key int, el N, parent *avlNode[N]) *avlNode[N] {
	if a == nil {
		return &avlNode[N]{
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

func (a *avlNode[N]) delete(key int) *avlNode[N] {
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

func (a *avlNode[N]) find(k int) (*avlNode[N], bool) {
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

func (a *avlNode[N]) rebalance() *avlNode[N] {
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

func (a *avlNode[N]) rebalanceLeft() *avlNode[N] {
	l := a.left
	if l.left.getHeight() < l.right.getHeight() {
		// LR imbalance
		a.left = l.rotateLeft()
	}
	// LL imbalance
	return a.rotateRight()

}

func (a *avlNode[N]) rebalanceRight() *avlNode[N] {
	r := a.right
	if r.right.getHeight() < r.left.getHeight() {
		// RL imbalance
		a.right = r.rotateRight()
	}
	// RR imbalance
	return a.rotateLeft()
}

func (a *avlNode[N]) rotateLeft() *avlNode[N] {
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

func (a *avlNode[N]) rotateRight() *avlNode[N] {
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

func (a *avlNode[N]) getHeight() int {
	if a == nil {
		return 0
	}
	return a.height
}

func (a *avlNode[N]) updateHeight() {
	if a == nil {
		return
	}
	maxHeight := a.left.getHeight()
	if maxHeight < a.right.getHeight() {
		maxHeight = a.right.getHeight()
	}
	a.height = 1 + maxHeight

}

func (a *avlNode[N]) isLeaf() bool {
	if a == nil {
		return true
	}
	return a.left == nil && a.right == nil
}

func (a *avlNode[N]) leftDescendant() *avlNode[N] {
	if a.left == nil {
		return a
	}
	return a.left.leftDescendant()
}

func (a *avlNode[N]) rightAncestor() *avlNode[N] {
	if a == nil || a.parent == nil {
		return nil
	}
	if a.key < a.parent.key {
		return a.parent
	}
	return a.parent.rightAncestor()
}

func (a *avlNode[N]) rangeSearch(low int, high int) []N {
	return []N{}
}

func (a *avlNode[N]) next() *avlNode[N] {
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

type AVLTree[N any] struct {
	root *avlNode[N]
}

func NewAVLTree[N any]() *AVLTree[N] {
	return &AVLTree[N]{}
}

func (t *AVLTree[N]) Insert(k int, el N) {
	t.root = t.root.insert(k, el, t.root)
}

func (t *AVLTree[N]) Find(key int) (N, error) {
	n, ok := t.root.find(key)
	if !ok {
		var z N
		return z, errors.New("not found")
	}
	return n.data, nil
}

func (t *AVLTree[N]) RangeSearch(low, high int) []N {
	result := make([]N, 0)
	n, _ := t.root.find(low)
	for n != nil && n.key <= high {
		if n.key >= low {
			result = append(result, n.data)
		}
		n = n.next()
	}
	return result
}

func (t AVLTree[N]) Traverse(strategy Strategy) []N {
	result := make([]N, 0)
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

func (t *AVLTree[N]) Delete(key int) {
	t.root = t.root.delete(key)
}

func traversePostOrder[N any](root *avlNode[N], result *[]N) {
	if root.left != nil {
		traversePostOrder(root.left, result)
	}

	if root.right != nil {
		traversePostOrder(root.right, result)
	}
	*result = append(*result, root.data)
}

func traversePreOrder[N any](root *avlNode[N], result *[]N) {
	*result = append(*result, root.data)

	if root.left != nil {
		traversePreOrder(root.left, result)
	}

	if root.right != nil {
		traversePreOrder(root.right, result)
	}
}

func traverseInOrder[N any](root *avlNode[N], result *[]N) {
	if root.left != nil {
		traverseInOrder(root.left, result)
	}
	*result = append(*result, root.data)
	if root.right != nil {
		traverseInOrder(root.right, result)
	}
}
