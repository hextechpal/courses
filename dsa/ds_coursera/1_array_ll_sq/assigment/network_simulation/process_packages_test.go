package network_simulation

import (
	"reflect"
	"testing"
)

func TestProcessPackets(t *testing.T) {
	type args struct {
		size     int
		requests []Request
	}
	tests := []struct {
		name string
		args args
		want []Response
	}{
		{
			name: "test0",
			args: args{
				size: 1,
				requests: []Request{{0, 1}, {0,1}},
			},
			want: []Response{{false, 0}, {true, -1}},
		},
		{
			name: "test0",
			args: args{
				size: 3,
				requests: []Request{{0, 1}, {3,1}, {10, 1}},
			},
			want: []Response{{false, 0}, {false, 3}, {false, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessPackets(tt.args.size, tt.args.requests); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessPackets() = %v, want %v", got, tt.want)
			}
		})
	}
}
