package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	M1 = 1000000007
	M2 = 1000000009
	X  = 237
)

func substringMatch(s string, i1, i2, l int) bool {
	h1, h2 := preCompute(s)

	val1 := hashValue(h1, M1, i1, l)
	val2 := hashValue(h2, M2, i1, l)

	val3 := hashValue(h1, M1, i2, l)
	val4 := hashValue(h2, M2, i2, l)

	if val1 == val3 && val2 == val4 {
		return true
	}
	return false
}

func hashValue(hashes []float64, m float64, i, l int) float64 {
	mod := math.Mod(math.Pow(X, float64(l)), m)
	return math.Mod(hashes[i+l], m) - math.Mod(mod*math.Mod(hashes[i], m), m)
}

func preCompute(s string) ([]float64, []float64) {
	h1 := make([]float64, len(s)+1)
	h2 := make([]float64, len(s)+1)
	h1[0] = 0
	h2[0] = 0
	for i := 1; i < len(s)+1; i++ {
		h1[i] = math.Mod(X*h1[i-1]+float64(s[i-1]), M1)
		h2[i] = math.Mod(X*h2[i-1]+float64(s[i-1]), M2)
	}
	return h1, h2
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l1, _ := r.ReadString('\n')
	s := strings.Trim(l1, "\n")

	l2, _ := r.ReadString('\n')
	ops, _ := strconv.Atoi(strings.Trim(l2, "\n"))

	cmds := make([]string, ops)
	for i := 0; i < ops; i++ {
		l2, _ := r.ReadString('\n')
		cmds[i] = strings.Trim(l2, "\n")
	}

	for _, cmd := range cmds {
		arr := strings.Split(cmd, " ")
		i1, _ := strconv.Atoi(arr[0])
		i2, _ := strconv.Atoi(arr[1])
		l, _ := strconv.Atoi(arr[2])

		fmt.Printf("Comparing %s | %s : ", s[i1:i1+l], s[i2:i2+l])
		if substringMatch(s, i1, i2, l) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
