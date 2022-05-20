package checkbrackets

import (
	"github.com/hextechpal/dsa/lib/stack"
)

var openBrackets = map[rune]bool{
	'{': true,
	'[': true,
	'(': true,
}

var closeBrackets = map[rune]bool{
	'}': true,
	']': true,
	')': true,
}

var pairs = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

func check(in string) int {
	s := stack.NewStack(len(in))
	for i, r := range []rune(in) {
		if _, ok := openBrackets[r]; ok {
			s.Push(r)
		} else if _, ok := closeBrackets[r]; ok {
			e, err := s.Pop()
			if err != nil || r != pairs[e.(rune)] {
				return i + 1
			}
		}
	}

	if !s.Empty() {
		return s.Size()
	}
	return -1
}
