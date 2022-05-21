package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	p = 1000000007
	x = 263
)

type Matcher struct {
	Pattern string
	Text    string
}

func (m *Matcher) find() []int {
	result := make([]int, 0)
	phash := m.hash(m.Pattern)
	preHashes := m.preComputeHashes()

	for i := 0; i <= len(m.Text)-len(m.Pattern); i++ {
		if preHashes[i] != phash {
			continue
		}

		if m.Text[i:i+len(m.Pattern)] == m.Pattern {
			result = append(result, i)
		}
	}
	return result
}

func (m *Matcher) hash(s string) float64 {
	hash := float64(0)
	for i := len(s) - 1; i >= 0; i-- {
		hash = math.Mod(hash*x+float64(s[i]), p)
	}
	return hash
}

func (m *Matcher) preComputeHashes() []float64 {
	window := len(m.Text) - len(m.Pattern)
	hashes := make([]float64, window+1)
	s := m.Text[window:]
	hashes[window] = m.hash(s)

	y := float64(1)
	for i := 1; i <= len(m.Pattern); i++ {
		y = math.Mod(y*x, p)
	}

	for i := window - 1; i >= 0; i-- {
		v := x*hashes[i+1] + float64(m.Text[i]) - math.Mod(y*float64(m.Text[i+len(m.Pattern)]), p)
		hashes[i] = math.Mod(v+p, p)
	}
	return hashes
}

func main() {
	r := bufio.NewReader(os.Stdin)
	pattern, _ := r.ReadString('\n')
	text, _ := r.ReadString('\n')

	m := &Matcher{
		Pattern: strings.Trim(pattern, "\n"),
		Text:    strings.Trim(text, "\n"),
	}
	for _, idx := range m.find() {
		fmt.Printf("%d ", idx)
	}
}
