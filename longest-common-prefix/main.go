package main

import "fmt"

func main() {
	sliceStr := []string{"flower", "flow", "flight"}
	commonPrefix := findCommonPrefix(sliceStr)
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
