package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "AAAFFFFVCCDDDH"
	minStr := createMinStr(str)
	fmt.Println(minStr)
}

func createMinStr(s string) string {
	var count int
	var lastRune rune
	sliceRune := make([]string, 0, len(s))
	for _, val := range s {
		if lastRune == val {
			count++
			continue
		}
		sliceRune = addRune(sliceRune, count, lastRune)
		lastRune = val
		count = 1
	}
	sliceRune = addRune(sliceRune, count, lastRune)
	return strings.Join(sliceRune, "")
}

func addRune(sliceRune []string, count int, lastRune rune) []string {
	if count < 2 {
		sliceRune = append(sliceRune, string(lastRune))
	}
	if count > 1 {
		sliceRune = append(sliceRune, string(lastRune), strconv.Itoa(count))
	}
	return sliceRune
}
