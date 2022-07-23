package assignment

import "fmt"

type opInfo struct {
	count int
	op    string
}

func countOps(n int) int {
	opc := make([]opInfo, n+1)
	opc[1] = opInfo{count: 0, op: ""}
	for i := 2; i < n+1; i++ {
		min := opc[i-1].count + 1
		op := "add"
		if i%2 == 0 && min > opc[i/2].count+1 {
			min = opc[i/2].count + 1
			op = "mul"
		}
		if i%3 == 0 && min > opc[i/3].count+1 {
			min = opc[i/3].count + 1
			op = "div"
		}
		opc[i] = opInfo{
			count: min,
			op:    op,
		}
	}

	i := n
	str := fmt.Sprintf("%d", n)
	for i != 1 {
		op := opc[i].op
		if op == "add" {
			i = i - 1
		} else if op == "mul" {
			i = i / 2
		} else if op == "div" {
			i = i / 3
		}
		str = fmt.Sprintf("%d %s", i, str)
	}
	fmt.Println(str)
	return opc[n].count
}
