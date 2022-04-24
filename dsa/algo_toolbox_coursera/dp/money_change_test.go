package dp

import "testing"

func Test_changeMoney(t *testing.T) {
	type args struct {
		money int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Change 2", args{money: 2, coins: []int{1,3,4}}, 2},
		{"Change 34", args{money: 34, coins: []int{1,3,4}}, 9},
		{"Change 33", args{money: 33, coins: []int{2}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeMoney(tt.args.money, tt.args.coins); got != tt.want {
				t.Errorf("changeMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
