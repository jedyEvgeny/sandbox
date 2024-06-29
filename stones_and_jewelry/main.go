package main

import (
	"fmt"
)

func main() {
	var J, S string
	fmt.Scan(&J, &S)
	countJ := findJ(J, S)
	fmt.Println(countJ)
}

func findJ(j, s string) int {
	jRune := []rune(j)
	sRune := []rune(s)
	var count int
	for i := 0; i < len(sRune); i++ {
		for k := 0; k < len(jRune); k++ {
			if sRune[i] == jRune[k] {
				count++
				break
			}
		}
	}
	return count
}
