package assignment

import "testing"

func Test_editDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{"ab", "ab"}, 0},
		{"Test2", args{"short", "ports"}, 3},
		{"Test3", args{"editing", "distance"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := editDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("editDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
