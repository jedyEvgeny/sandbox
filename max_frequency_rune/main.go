package main

import (
	"fmt"
)

func main() {
	str := ""
	maxFrequencyRune := findMaxFrequencyRune(str)
	fmt.Println(string(maxFrequencyRune))
}

func findMaxFrequencyRune(s string) rune {
	mapFrequenceRune := make(map[rune]int)
	var ans rune
	var count int
	for _, el := range s {
		mapFrequenceRune[el]++
		if mapFrequenceRune[el] > count {
			count = mapFrequenceRune[el]
			ans = el
		}
	}
	return ans
}
