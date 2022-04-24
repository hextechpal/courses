package checkbrackets

import "testing"

func Test_check_files(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bb_0",
			args: args{"{}"},
			want: -1,
		},
		{
			name: "bb_1",
			args: args{"{}[]"},
			want: -1,
		},
		{
			name: "bb_2",
			args: args{"[()]"},
			want: -1,
		},
		{
			name: "bb_3",
			args: args{"(())"},
			want: -1,
		},
		{
			name: "bb_4",
			args: args{"{[]}()"},
			want: -1,
		},
		{
			name: "bb_5",
			args: args{"foo(bar);"},
			want: -1,
		},
		{
			name: "unb_1",
			args: args{"{"},
			want: 1,
		},
		{
			name: "unb_2",
			args: args{"{[}"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.in); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}