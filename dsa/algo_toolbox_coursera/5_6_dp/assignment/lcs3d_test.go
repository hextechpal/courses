package assignment

import "testing"

func Test_lcs3d(t *testing.T) {
	type args struct {
		seq1 []int
		seq2 []int
		seq3 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{seq1: []int{1, 2, 3}, seq2: []int{2, 1, 3}, seq3: []int{1, 3, 5}}, 2},
		{"Test1", args{seq1: []int{8, 3, 2, 1, 7}, seq2: []int{8, 2, 1, 3, 8, 10, 7}, seq3: []int{6, 8, 3, 1, 4, 7}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcs3d(tt.args.seq1, tt.args.seq2, tt.args.seq3); got != tt.want {
				t.Errorf("lcs3d() = %v, want %v", got, tt.want)
			}
		})
	}
}
