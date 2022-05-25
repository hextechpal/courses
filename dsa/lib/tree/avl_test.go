package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAvlNode_Find(t *testing.T) {
	tree := createAvlTree(t)

	type args struct {
		key int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		found bool
	}{
		{"find Root", args{key: 7}, 7, true},
		{"find Leaf", args{key: 20}, 20, true},
		{"find Leaf", args{key: 200}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := tree.find(tt.args.key)

			if found != tt.found {
				t.Errorf("tree.find() error = %v, found %v", found, tt.found)
				return
			}

			if found && !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvlNode_Insert(t *testing.T) {
	tr := NewAVLTree[int]()
	tr.Insert(9, 9)
	tr.Insert(6, 6)
	tr.Insert(7, 7)
	tr.Insert(7, 7)

	fmt.Println(tr.RangeSearch(5, 8))

	tr.Delete(10)
	tr.Delete(5)
	tr.Delete(2)

	fmt.Println(tr.Traverse(INORDER))
}

func createAvlTree(t *testing.T) *avlNode[int] {
	t.Helper()
	n9 := &avlNode[int]{key: 9, data: 9}
	n6 := &avlNode[int]{key: 6, data: 6}
	n7 := &avlNode[int]{key: 7, data: 7}
	n2 := &avlNode[int]{key: 2, data: 2}
	n5 := &avlNode[int]{key: 5, data: 5}
	n10 := &avlNode[int]{key: 10, data: 10}
	n20 := &avlNode[int]{key: 20, data: 20}

	n7.left = n5
	n7.right = n10

	n5.left = n2
	n5.right = n6

	n10.left = n9
	n10.right = n20

	return n7
}
