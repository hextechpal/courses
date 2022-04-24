package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hextechpal/dsa/ds_coursera/6_splay/assignment/splay"
)

const M = 1000000001

func main() {
	reader := bufio.NewReader(os.Stdin)
	l1, _ := reader.ReadString('\n')
	c := strings.TrimSpace(l1)
	opc, _ := strconv.Atoi(c)
	tree := splay.NewTree()
	sum := 0
	for i := 0; i < opc; i++ {
		l2, _ := reader.ReadString('\n')
		in := strings.Split(strings.TrimSpace(l2), " ")
		op := in[0]
		switch op {
		case "+":
			n, _ := strconv.Atoi(in[1])
			tree.Insert((n + sum) % M)
		case "-":
			n, _ := strconv.Atoi(in[1])
			tree.Delete((n + sum) % M)
		case "?":
			n, _ := strconv.Atoi(in[1])
			found := tree.Find((n + sum) % M)
			if found {
				fmt.Println("Found")
			} else {
				fmt.Println("Not Found")
			}
		case "s":
			l, _ := strconv.Atoi(in[1])
			r, _ := strconv.Atoi(in[2])
			sum = tree.Sum((l+sum)%M, (r+sum)%M)
			fmt.Println(sum)
		}
	}
}
