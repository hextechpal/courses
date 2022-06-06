package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildTrie(patterns []string) []map[rune]int {
	trie := make([]map[rune]int, 0)
	trie = append(trie, make(map[rune]int))
	for _, pattern := range patterns {
		current := 0
		for _, c := range pattern {
			j, ok := trie[current][c]
			if ok {
				current = j
			} else {
				trie = append(trie, make(map[rune]int))
				trie[current][c] = len(trie) - 1
				current = len(trie) - 1
			}
		}
	}
	return trie
}

func printTrie(trie []map[rune]int) {
	for i := 0; i < len(trie); i++ {
		for w, v := range trie[i] {
			fmt.Printf("%d->%d:%s\n", i, v, string(w))
		}

	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l1, _ := r.ReadString('\n')

	pc, _ := strconv.Atoi(strings.Trim(l1, "\n"))
	patterns := make([]string, pc)
	for i := 0; i < pc; i++ {
		l2, _ := r.ReadString('\n')
		patterns[i] = strings.Trim(l2, "\n")
	}
	printTrie(buildTrie(patterns))
}
