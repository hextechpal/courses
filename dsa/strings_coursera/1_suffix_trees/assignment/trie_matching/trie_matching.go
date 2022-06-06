package main

import "fmt"

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

func match(text string, patterns []string) []int {
	ans  := make([]int, 0)
	trie := buildTrie(patterns)
	fmt.Printf("%v\n", trie)
	h := 0
	for h < len(text) {
		fmt.Printf("prefixTrieMatch trie match %s\n", text[h:])
		if prefixTrieMatch(text[h:], trie){
			ans = append(ans, h)
		}
		h++
	}
	return ans

}

func prefixTrieMatch(s string, trie []map[rune]int) bool {
	f := 0
	v := 0

	for {
		// This is a leaf
		if len(trie[v]) == 0 {
			return true
		} else {
			if len(s) <= f {
				return false
			}else if w, ok := trie[v][rune(s[f])]; ok{
				f++
				v = w
			}else{
				return false
			}
		}
	}
}

func main(){
	fmt.Println(match("AAA", []string{"AA"}))
	fmt.Println()
	fmt.Println(match("AATCGGGTTCAATCGGGGT", []string{"ATCG", "GGGT"}))
}