package bst

import (
	"reflect"
	"testing"
)

type Integer struct {
	val int
}

func (i Integer) GetValue() int {
	return i.val
}
func (i Integer) Compare(c Valuer) int {
	if i.val < c.GetValue() {
		return -1
	} else if i.val == c.GetValue() {
		return 0
	} else {
		return 1
	}
}

func TestAvlTree_Find(t *testing.T) {
	tree := createAvlTree(t)
	v1 := tree.Search(Integer{20})
	if v1.GetValue() != 20 {
		t.Errorf("Wanted 20 got %d\n", v1.GetValue())
		t.Fail()
	}

	v2 := tree.Search(Integer{100})
	if v2 != nil {
		t.Errorf("Wanted nil got %d\n", v2.GetValue())
		t.Fail()
	}
	tree.InOrder()
}

func TestAvlTree_AvlInsert(t *testing.T) {
	tree := &AvlTree{}
	tree.AvlInsert(Integer{10})
	//tree.AvlInsert(Integer{7})
	//tree.AvlInsert(Integer{5})
	//tree.AvlInsert(Integer{2})
	tree.AvlInsert(Integer{30})
	tree.AvlInsert(Integer{20})
	//tree.AvlInsert(Integer{9})
	//tree.AvlInsert(Integer{6})
	got := tree.InOrder()
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("tree elements are not equal want %v got %v", want, got)
		t.Fail()
	}
}

func createAvlTree(t *testing.T) *AvlTree {
	t.Helper()
	n9 := &AvlNode{key: Integer{9}}
	n6 := &AvlNode{key: Integer{6}}
	n7 := &AvlNode{key: Integer{7}}
	n2 := &AvlNode{key: Integer{2}}
	n5 := &AvlNode{key: Integer{5}}
	n10 := &AvlNode{key: Integer{10}}
	n20 := &AvlNode{key: Integer{20}}

	n10.left = n5
	n10.right = n20

	n5.parent = n10
	n5.left = n2
	n5.right = n7

	n2.parent = n5

	n7.parent = n5
	n7.left = n6
	n7.right = n9

	n6.parent = n7
	n9.parent = n7

	n20.parent = n10

	return &AvlTree{root: n10}
}
