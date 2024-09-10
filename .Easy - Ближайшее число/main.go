/*
Кодран - https://coderun.yandex.ru/selections/quickstart/problems/nearest-number/description?currentPage=1&pageSize=20&search=
Уровень easy

Напишите программу, которая находит в массиве элемент, самый близкий по величине к данному числу.

Формат ввода
В первой строке задается одно натуральное число N, не превосходящее 1000 — размер массива. Во второй строке содержатся 
N чисел — элементы массива, целые числа, не превосходящие по модулю 1000. В третьей строке вводится одно целое число 
x, не превосходящее по модулю 1000.

Формат вывода
Вывести значение элемента массива, ближайшее к 
x. Если таких чисел несколько, выведите любое из них.

ЗЫ решил с первой попытки!
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var length int
	_, _ = fmt.Scan(&length)
	slice := make([]float64, length)
	for i := 0; i < length; i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	var target float64
	_, _ = fmt.Scan(&target)
	closerTarget := slice[0]
	lastDiff := math.Sqrt((target - closerTarget) * (target - closerTarget))
	if lastDiff == 0 {
		fmt.Println(closerTarget)
		return
	}
	var currDiff float64
	for i := 1; i < length; i++ {
		currDiff = math.Sqrt((slice[i] - target) * (slice[i] - target))
		if currDiff == 0 {
			closerTarget = slice[i]
			break
		}
		if currDiff < lastDiff {
			closerTarget = slice[i]
		}
		lastDiff = currDiff
	}
	fmt.Println(closerTarget)
}
