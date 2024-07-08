package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceStr := []string{"flower", "flow", "flight"}
	commonPrefix := findCommonPrefix(sliceStr)
	// commonPrefix := longestCommonPrefix(sliceStr)
	// commonPrefix := longestCommonPrefixShort(sliceStr)
	fmt.Println(commonPrefix)
}

func findCommonPrefix(arr []string) string {
	var (
		ans               string
		minLenStr         int
		lastLetter        rune
		countCommonLetter int
	)
	for idx, el := range arr {
		if idx == 0 {
			val := []rune(el)
			minLenStr = len(val)
		}
		if len([]rune(el)) < minLenStr {
			minLenStr = len([]rune(el))
		}
	}
	// fmt.Println(minLenStr)
	for i := 0; i < minLenStr; i++ {
		var val []rune
		for idx, el := range arr {
			val = []rune(el)
			if idx == 0 {
				lastLetter = val[i]
				// fmt.Println(string(lastLetter))
			}
			if val[i] == lastLetter {
				countCommonLetter++
			}
		}
		if countCommonLetter == len(arr) {
			ans += string(lastLetter)
			// fmt.Println("---", string(ans))
		}
		if countCommonLetter != len(arr) {
			break
		}
		countCommonLetter = 0
	}
	return ans
}

// Более короткий и эффективный алгоритм
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// Ещё более кототкий алгоритм
func longestCommonPrefixShort(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}
