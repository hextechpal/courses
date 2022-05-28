package main

import (
	"fmt"
	"math"
)

const P1 = 100003
const P2 = 100019
const X = 31

func lcs(s1, s2 string) (int, int, int) {
	var ss string
	var ls string

	if len(s1) < len(s2) {
		ss, ls = s1, s2
	} else {
		ss, ls = s2, s1
	}

	return maxlength(ls, ss, 0, len(ss), 0, 0, 0)
}

func maxlength(ls string, ss string, low int, high int, maxLen int, lstart int, sstart int) (int, int, int) {
	if low > high {
		return lstart, sstart, maxLen
	}
	mid := (low + high) / 2
	hl1 := hashTable(ls, mid, P1)
	hl2 := hashTable(ls, mid, P2)

	hm1 := hashMap(ss, mid, P1)
	hm2 := hashMap(ss, mid, P2)

	pos1, check1 := match(hl1, hm1)
	pos2, check2 := match(hl2, hm2)

	if check1 && check2 {
		for lsi, ssi := range pos1 {
			if _, ok := pos2[lsi]; ok {
				return maxlength(ls, ss, mid+1, high, mid, lsi, ssi)
			}
		}
	}

	return maxlength(ls, ss, low, mid-1, maxLen, lstart, sstart)
}

func match(ht []float64, hm map[float64]int) (map[int]int, bool) {
	positions := make(map[int]int)
	found := false
	for lstart, h := range ht {
		if sstart, ok := hm[h]; ok {
			positions[lstart] = sstart
			found = true
		}
	}
	return positions, found
}

func hashTable(s string, pLen int, p float64) []float64 {
	window := len(s) - pLen
	hashes := make([]float64, window+1)
	subStr := s[window:]
	hashes[window] = hash(subStr, p)

	y := float64(1)
	for i := 0; i < pLen; i++ {
		y = math.Mod(y*X, p)
	}

	for i := window - 1; i >= 0; i-- {
		v := X*hashes[i+1] + float64(s[i]) - math.Mod(y*float64(s[i+pLen]), p)
		hashes[i] = math.Mod(v+p, p)
	}
	return hashes
}

func hashMap(s string, pLen int, p float64) map[float64]int {
	hm := make(map[float64]int)
	window := len(s) - pLen
	subStr := s[window:]
	last := hash(subStr, p)
	hm[last] = window

	y := float64(1)
	for i := 0; i < pLen; i++ {
		y = math.Mod(y*X, p)
	}

	for i := window - 1; i >= 0; i-- {
		v := X*last + float64(s[i]) - math.Mod(y*float64(s[i+pLen]), p)
		last = math.Mod(v+p, p)
		hm[last] = i
	}
	return hm
}

func hash(s string, p float64) float64 {
	h := float64(0)
	for i := len(s) - 1; i >= 0; i-- {
		h = math.Mod(h*X+float64(s[i]), p)
	}
	return h
}

func main() {
	fmt.Println(lcs("aabaa", "babbaab"))
}
