package assignment

import "testing"

func Test_longestCommonSubSequence(t *testing.T) {
	type args struct {
		seq1 []int
		seq2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{[]int{2, 7, 5}, []int{2, 5}}, 2},
		{"Test2", args{[]int{7}, []int{1, 2, 3, 4}}, 0},
		{"Test3", args{[]int{2, 7, 8, 3}, []int{5, 2, 8, 7}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubSequence(tt.args.seq1, tt.args.seq2); got != tt.want {
				t.Errorf("longestCommonSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
