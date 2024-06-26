package main

import (
	"fmt"
)

func main() {
	fmt.Println("Римские цифры в арабские")
	fmt.Println("Введите римские цифры в формате I, V, X, L, C, D, M")
	var str string
	fmt.Scan(&str)
	num := romanToInt(str)
	fmt.Printf("%v = %d", str, num)
}

func romanToInt(s string) int {
	repo := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sliceRoman := make([]string, len(s))
	for idx, v := range s {
		sliceRoman[idx] = string(v)
	}
	var result int
	for i := 0; i < len(s); i++ {
		valNum := repo[sliceRoman[i]]
		result = result + valNum
		if i > 0 {
			if (sliceRoman[i] == "V" || sliceRoman[i] == "X") && sliceRoman[i-1] == "I" {
				result--
				if i < len(s)-1 {
					i++
				}
			}
			if (sliceRoman[i] == "L" || sliceRoman[i] == "C") && sliceRoman[i-1] == "X" {
				result -= 10
				if i < len(s)-1 {
					i++
				}
			}
			if (sliceRoman[i] == "D" || sliceRoman[i] == "M") && sliceRoman[i-1] == "C" {
				result -= 100
				if i < len(s)-1 {
					i++
				}
			}
		}
	}
	return result
}
