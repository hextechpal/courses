package splay
//
//import "fmt"
//
//type Tree struct {
//	root *Node
//}
//
//func (t *Tree) Find(k int) bool {
//	t.root = t.root.Find(k)
//	return t.root != nil && t.root.key == k
//}
//
//func (t *Tree) Insert(k int) {
//	t.root = splay(t.root, k)
//	if t.root != nil && t.root.key == k {
//		return
//	}
//	t.root = t.root.Insert(k)
//}
//
//func (t *Tree) Delete(k int) {
//	t.root = splay(t.root, k)
//	if t.root == nil || t.root.key != k {
//		return
//	}
//	t.root = t.root.Delete()
//}
//
//func (t *Tree) InOrder() {
//	t.root.inOrder()
//}
//
//func (t *Tree) Sum(l, r int)  int{
//	ans := 0
//	left, mid := t.root.Split(l)
//	mid, right := mid.Split(r+1)
//
//	if mid == nil{
//		t.root = Merge(left, right)
//	}else{
//		ans = mid.sum
//		t.root = Merge(left, Merge(mid, right))
//	}
//	return ans
//}
//
//
//func NewTree() *Tree {
//	return &Tree{}
//}
//
//type Node struct {
//	key int
//	sum int
//
//	left  *Node
//	right *Node
//}
//
//func (n *Node) IsLeaf() bool {
//	return n.left == nil && n.right == nil
//}
//
//func (n *Node) Find(k int) *Node {
//	return splay(n, k)
//}
//
//func (n *Node) Split(k int) (*Node, *Node){
//	root := splay(n, k)
//	if root == nil{
//		return nil, nil
//	}
//	if root.key < k {
//		r1 := root.right
//		root.right = nil
//		root.UpdateSum()
//		return root, r1
//	}
//
//	r1 := root.left
//	root.left = nil
//	root.UpdateSum()
//	return r1, root
//}
//
//func splay(root *Node, k int) *Node {
//	if root == nil || root.key == k {
//		return root
//	}
//
//	// Search left subtree
//	if root.key > k {
//		if root.left == nil {
//			return root
//		}
//
//		if root.left.key > k {
//			// zig-zig formation
//			root.left.left = splay(root.left.left, k)
//			root = rotateRight(root)
//		} else if root.left.key < k {
//			// zag-zig formation
//			root.left.right = splay(root.left.right, k)
//			if root.left.right != nil {
//				root.left = rotateLeft(root.left)
//			}
//		}
//		if root.left != nil {
//			root = rotateRight(root)
//		}
//	} else if root.key < k {
//		if root.right == nil {
//			return root
//		}
//		if root.right.key < k {
//			//zag-zag formation
//			root.right.right = splay(root.right.right, k)
//			root = rotateLeft(root)
//		} else if root.right.key > k {
//			//zig-zag formation
//			root.right.left = splay(root.right.left, k)
//			if root.right.left != nil {
//				root.right = rotateRight(root.right)
//			}
//		}
//		if root.right != nil {
//			root = rotateLeft(root)
//		}
//	}
//	return root
//}
//
//func rotateLeft(n *Node) *Node {
//	r := n.right
//	n.right = r.left
//	r.left = n
//	n.UpdateSum()
//	r.UpdateSum()
//	return r
//}
//
//func rotateRight(n *Node) *Node {
//	l := n.left
//	n.left = l.right
//	l.right = n
//	n.UpdateSum()
//	l.UpdateSum()
//	return l
//}
//
//func (n *Node) Insert(k int) *Node {
//	newNode := &Node{
//		key:   k,
//		sum:   k,
//		left:  nil,
//		right: nil,
//	}
//	if n == nil {
//		return newNode
//	}
//
//	if n.key < k {
//		newNode.left = n
//		newNode.right = n.right
//		n.right = nil
//	} else if n.key > k {
//		newNode.left = n.left
//		newNode.right = n
//		n.left = nil
//	}
//	n.UpdateSum()
//	newNode.UpdateSum()
//	return newNode
//}
//
//func (n *Node) search(k int) *Node {
//	if n == nil {
//		return nil
//	}
//
//	if n.key == k {
//		return n
//	} else if n.key < k {
//		return n.right.search(k)
//	} else {
//		return n.left.search(k)
//	}
//}
//
//func (n *Node) inOrder() {
//	if n == nil {
//		return
//	}
//	if n.left != nil {
//		n.left.inOrder()
//	}
//	fmt.Printf("%d [%d]", n.key, n.sum)
//	if n.right != nil {
//		n.right.inOrder()
//	}
//}
//
//func (n *Node) Delete() *Node {
//	root1 := n.left
//	root2 := n.right
//	if root1  == nil{
//		return root2
//	}
//	max := root1.FindMax()
//	root1 = splay(root1 , max.key)
//	root1.right = root2
//	root1.UpdateSum()
//	return root1
//}
//
//func (n *Node) FindMax() *Node {
//	it := n
//	for n.right != nil {
//		it = n.right
//	}
//	return it
//}
//
//func (n *Node) Sum() int {
//	if n == nil {
//		return 0
//	}
//	return n.sum
//}
//
//func (n *Node) UpdateSum() {
//	n.sum = n.right.Sum() + n.left.Sum() + n.key
//}
//
//func Merge(r1 *Node, r2 *Node) *Node{
//	if r1 == nil {
//		return r2
//	}
//	max := r1.FindMax()
//	root := splay(r1, max.key)
//	root.right = r2
//	root.UpdateSum()
//	return root
//}