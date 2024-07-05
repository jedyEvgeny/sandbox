package main

import "fmt"

func main() {
	var strFirst, strSecond string
	fmt.Scan(&strFirst, &strSecond)
	if len(strFirst) != len(strSecond) {
		fmt.Println(0)
		return
	}
	mapRuneFirstStr := make(map[rune]int)
	for idx, el := range strFirst {
		mapRuneFirstStr[el]++
		valSec := rune(strSecond[idx])
		mapRuneFirstStr[valSec]--
		if mapRuneFirstStr[el] == 0 {
			delete(mapRuneFirstStr, el)
		}
		if mapRuneFirstStr[valSec] == 0 {
			delete(mapRuneFirstStr, valSec)
		}
	}
	if len(mapRuneFirstStr) == 0 {
		fmt.Println(1)
	}
	if len(mapRuneFirstStr) != 0 {
		fmt.Println(0)
	}
}
