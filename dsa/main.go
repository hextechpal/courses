package main

import (
	"fmt"
	"github.com/hextechpal/dsa/lib/sort/merge"
	"github.com/hextechpal/dsa/lib/sort/quick"
)

func main() {
	fmt.Println(quick.Sort([]int{13, 11, 7, 4, 12}))

	fmt.Println(merge.Sort([]int{11, 13, 7, 4, 19}))
}
