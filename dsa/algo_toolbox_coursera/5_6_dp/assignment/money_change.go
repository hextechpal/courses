package assignment

import (
	"fmt"
	"math"
)

func changeMoney(money int, coins []int) int {
	cc := make([]int, money+1)
	for i := 1; i < money+1; i++ {
		cc[i] = math.MaxInt
	}
	for m := 0; m <= money; m++ {
		for _, coin := range coins {
			cr := 0
			if m-coin >= 0 && cc[m-coin] != -1 {
				cr = cc[m-coin] + 1
				if cc[m] > cr {
					cc[m] = cr
				}
			}
			if cc[m] == math.MaxInt {
				cc[m] = -1
			}
		}
	}
	fmt.Println(cc)
	return cc[money]
}
