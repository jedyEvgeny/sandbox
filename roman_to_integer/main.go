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

	var result int
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		val := repo[string(s[i])]
		if val < prev {
			result -= val
		} else {
			result += val
		}
		prev = val
	}

	return result
}
