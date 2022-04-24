package treeheight

import "testing"

func TestTree_Height(t *testing.T) {
	type fields struct {
		n       int
		parents []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name : "test1",
			fields: fields{
				n : 5,
				parents: []int{4, -1, 4, 1, 1},
			},
			want: 3,
		},
		{
			name : "test1",
			fields: fields{
				n : 5,
				parents: []int{-1, 0, 4, 0, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				n:       tt.fields.n,
				parents: tt.fields.parents,
			}
			if got := tr.Height(); got != tt.want {
				t.Errorf("Tree.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
