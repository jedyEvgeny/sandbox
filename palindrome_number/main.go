package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var reversed int
	tempX := x
	for tempX != 0 {
		reversed = tempX%10 + reversed*10
		tempX /= 10
	}
	return reversed == x
}

func main() {
	fmt.Println("Поиск числа-палиндрома")
	fmt.Println("Палиндром - когда справа налево читается как слева направо")
	fmt.Println("-------------")
	fmt.Println("Введите число для проверки")
	var num int
	fmt.Scan(&num)
	ok := isPalindrome(num)
	fmt.Println(ok)
}
