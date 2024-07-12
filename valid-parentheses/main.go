package main

import (
	"fmt"
)

func main() {
	str := "(("
	ans := determineCorrectBrackets(str)
	fmt.Println(ans)
}

func determineCorrectBrackets(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	bracketSlice := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			bracketSlice = append(bracketSlice, r)
			continue
		}
		if len(bracketSlice) == 0 {
			return false
		}
		lastSimbolSliceRune := bracketSlice[len(bracketSlice)-1]
		if r == brackets[lastSimbolSliceRune] {
			bracketSlice = bracketSlice[:len(bracketSlice)-1]
			continue
		}
		return false
	}
	return len(bracketSlice) == 0
}
