package main

import (
	"reflect"
	"testing"
)

func Test_maxSliding(t *testing.T) {
	type args struct {
		nums  []int
		wsize int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test0",
			args: args{
				nums:  []int{2, 7, 3, 1, 5, 2, 6, 2},
				wsize: 4,
			},
			want: []int{7, 7, 5, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSliding(tt.args.nums, tt.args.wsize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSliding() = %v, want %v", got, tt.want)
			}
		})
	}
}
