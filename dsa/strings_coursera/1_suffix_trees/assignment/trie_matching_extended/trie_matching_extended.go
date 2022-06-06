package main

import "fmt"

type node struct {
	val        int
	endPattern bool
	edges      map[rune]*node
}

func (n *node) String() string {
	return fmt.Sprintf("val: %d, endPattern: %v, edges: %v", n.val, n.endPattern, n.edges)
}

func buildTrie(patterns []string) *node {
	root := &node{
		val:        0,
		endPattern: false,
		edges:      make(map[rune]*node),
	}
	nc := 0
	for _, pattern := range patterns {
		current := root
		for idx, c := range pattern {
			j, ok := current.edges[c]
			if ok {
				if !j.endPattern {
					fmt.Printf("Edge found for %v, isEnd, %v\n", c, idx == len(pattern)-1)
					j.endPattern = idx == len(pattern)-1
				}
				current = j
			} else {
				nc++
				n := &node{
					val:        nc,
					endPattern: idx == len(pattern)-1,
					edges:      make(map[rune]*node),
				}
				current.edges[c] = n
				current = n
			}
		}
	}
	return root
}

func match(text string, patterns []string) []int {
	ans := make([]int, 0)
	trie := buildTrie(patterns)
	fmt.Printf("%v\n", trie)
	h := 0
	for h < len(text) {
		if prefixTrieMatch(text[h:], trie) {
			ans = append(ans, h)
		}
		h++
	}
	return ans

}

func prefixTrieMatch(s string, trie *node) bool {
	f := 0
	root := trie

	for {
		// This is a leaf
		if root.endPattern {
			return true
		} else {
			if len(s) <= f {
				return false
			} else if w, ok := root.edges[rune(s[f])]; ok {
				f++
				root = w
			} else {
				return false
			}
		}
	}
}

func main() {
	fmt.Println(match("ACATA", []string{"AT", "A", "AG"}))
}
