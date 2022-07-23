package assignment

import "testing"

func Test_countOps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"OpC 1", args{1}, 1},
		{"OpC 5", args{5}, 3},
		{"OpC 96234", args{96234}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOps(tt.args.n); got != tt.want {
				t.Errorf("countOps() = %v, want %v", got, tt.want)
			}
		})
	}
}
