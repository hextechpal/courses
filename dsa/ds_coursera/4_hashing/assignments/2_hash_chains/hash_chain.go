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
	p = 1000000007
	x = 263
)

type hashTable struct {
	buckets int
	store   [][]string
}

func NewHashTable(buckets int) *hashTable {
	store := make([][]string, buckets)
	for i := 0; i < buckets; i++ {
		store[i] = make([]string, 0)
	}

	return &hashTable{
		buckets: buckets,
		store:   store,
	}
}

func (h *hashTable) Add(el string) {
	bucket := h.bucket(el)
	for _, s := range h.store[bucket] {
		if s == el {
			return
		}
	}
	h.store[bucket] = append(h.store[bucket], el)
}

func (h *hashTable) Del(el string) {
	bucket := h.bucket(el)
	for i, s := range h.store[bucket] {
		if s == el {
			h.store[bucket] = append(h.store[bucket][:i], h.store[bucket][i+1:]...)
		}
	}
}

func (h *hashTable) Find(el string) string {
	bucket := h.bucket(el)
	for _, s := range h.store[bucket] {
		if s == el {
			return "yes"
		}
	}
	return "no"
}

func (h *hashTable) Check(i int) {
	for j := len(h.store[i]) - 1; j >= 0; j-- {
		fmt.Printf("%s ", h.store[i][j])
	}
	fmt.Println()
}

func (h *hashTable) bucket(s string) int {
	return int(math.Mod(h.hash(s), float64(h.buckets)))
}

func (h *hashTable) hash(s string) float64 {
	hash := float64(0)
	for i := len(s) - 1; i >= 0; i-- {
		hash = math.Mod(hash*x+float64(s[i]), p)
	}
	return hash
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l1, _ := r.ReadString('\n')
	buckets, _ := strconv.Atoi(strings.Trim(l1, "\n"))

	h := NewHashTable(buckets)

	l2, _ := r.ReadString('\n')
	ops, _ := strconv.Atoi(strings.Trim(l2, "\n"))

	cmds := make([]string, ops)
	for i := 0; i < ops; i++ {
		op, _ := r.ReadString('\n')
		cmds[i] = op
	}

	for _, cmd := range cmds {
		args := strings.Split(strings.Trim(cmd, "\n"), " ")
		switch args[0] {
		case "add":
			h.Add(args[1])
		case "find":
			fmt.Println(h.Find(args[1]))
		case "del":
			h.Del(args[1])
		case "check":
			in, _ := strconv.Atoi(args[1])
			h.Check(in)
		}
	}
}
