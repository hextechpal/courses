package splay

type STree struct {
	root *SNode
}

func NewTree() *STree {
	return &STree{}
}

func (t *STree) Find(key int) bool {
	t.root = t.root.Find(key)
	if t.root != nil && t.root.key == key {
		return true
	}
	return false
}

func (t *STree) Insert(key int) {
	t.root = t.root.insert(key)
	_ = t.Find(key)
}

func (t *STree) Split(key int) (*SNode, *SNode) {
	return t.root.Split(key)
}

func (t *STree) Delete(key int) {
	next := t.root.next(key)
	if next != nil {
		t.root = splay(next)
	}
	found := t.Find(key)
	if !found {
		return
	}
	r := t.root.right
	if r == nil {
		t.root = t.root.left
		if t.root == nil {
			return
		}
		t.root.parent = nil
		updateSNode(t.root)
		return
	}

	r.left = t.root.left
	if t.root.left != nil {
		t.root.left.parent = r
	}
	t.root = r
	updateSNode(t.root)
}

func (t *STree) Sum(l int, r int) int {
	ans := 0
	left, mid := t.root.Split(l)
	mid, right := mid.Split(r + 1)

	if mid == nil {
		t.root = Merge(left, right)
	} else {
		ans = mid.sum
		t.root = Merge(left, Merge(mid, right))
	}
	return ans
}

type SNode struct {
	key    int
	sum    int
	left   *SNode
	right  *SNode
	parent *SNode
}

func (s *SNode) Sum() int {
	if s == nil {
		return 0
	}
	return s.sum
}

func (s *SNode) UpdateSum() {
	s.sum = s.right.Sum() + s.left.Sum() + s.key
}

func (s *SNode) Find(key int) *SNode {
	return splay(s.find(key))
}

func (s *SNode) find(key int) *SNode {
	if s == nil {
		return s
	}
	it := s
	for it.key != key {
		if it.left != nil && it.key > key {
			it = it.left
		} else if it.right != nil && it.key < key{
			it = it.right
		} else {
			break
		}
	}
	return it
}

func (s *SNode) insert(key int) *SNode {
	if s == nil {
		return &SNode{
			key:    key,
			sum:    key,
			left:   nil,
			right:  nil,
			parent: nil,
		}
	}
	if key < s.key {
		s.left = s.left.insert(key)
	} else {
		s.right = s.right.insert(key)
	}
	updateSNode(s)
	return s
}

func (s *SNode) Split(key int) (*SNode, *SNode) {
	root := s.Find(key)
	if root == nil {
		return nil, nil
	}
	if root.key < key {
		r1 := root.right
		root.right = nil
		root.UpdateSum()
		return root, r1
	}

	r1 := root.left
	root.left = nil
	root.UpdateSum()
	return r1, root

}

func (s *SNode) next(key int) *SNode {
	n := s.find(key)
	if n == nil {
		return nil
	}
	if n.right != nil {
		return LeftDecendant(n.right)
	} else {
		return RightAncestor(n)
	}
}

func RightAncestor(s *SNode) *SNode {
	for s.parent != nil && s.parent.right == s {
		s = s.parent
	}
	return s
}

func LeftDecendant(s *SNode) *SNode {
	for s.left != nil {
		s = s.left
	}
	return s
}

func splay(n *SNode) *SNode {
	if n == nil {
		return n
	}
	for n.parent != nil {
		if n.parent.parent == nil {
			smallRotation(n)
			break
		}
		bigRotation(n)
	}
	return n
}

func bigRotation(n *SNode) {
	if (n.parent.left == n && n.parent.parent.left == n.parent) || (n.parent.right == n && n.parent.parent.right == n.parent) {
		// zig zig : rotateRight and rotateRight
		smallRotation(n.parent)
		smallRotation(n)
	} else {
		// zig zag : rotateRight and rotateLeft
		smallRotation(n)
		smallRotation(n)
	}
}

func smallRotation(n *SNode) {
	p := n.parent
	if p == nil {
		return
	}
	gp := p.parent
	if p.left == n {
		m := n.right
		n.right = p
		p.left = m
	} else {
		m := n.left
		n.left = p
		p.right = m
	}
	updateSNode(p)
	updateSNode(n)

	n.parent = gp
	if gp != nil {
		if gp.left == p {
			gp.left = n
		} else {
			gp.right = n
		}
	}
}

func updateSNode(s *SNode) {
	if s == nil {
		return
	}
	s.UpdateSum()
	if s.left != nil {
		s.left.parent = s
	}
	if s.right != nil {
		s.right.parent = s
	}
}

func (s *SNode) FindMax() *SNode {
	it := s
	for s.right != nil {
		it = s.right
	}
	return it
}

func Merge(r1 *SNode, r2 *SNode) *SNode{
	if r1 == nil {
		return r2
	}
	max := r1.FindMax()
	root := splay(max)
	root.right = r2
	root.UpdateSum()
	return root
}
