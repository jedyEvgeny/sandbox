package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 1

	index := checkBinary(arr, target)

	if index != -1 {
		fmt.Printf("Элемент найден по индексу %d.\n", index)
	}
	if index == -1 {
		fmt.Println("Элемент не найден в массиве.")
	}
}

func checkBinary(arr []int, t int) int {
	var lowBound int            //Нижняя граница диапазона поиска изначально равна нулю
	highBound := len(arr) - 1   //Верхняя граница диапазона поиска
	ans := -1                   //Возвращаемое значение
	for lowBound <= highBound { //Пока эта часть не сократится до одного элемента
		mid := (lowBound + highBound) / 2 //Проверяем средний элемент с искомым значением
		if arr[mid] == t {                //Значение найдено
			ans = mid
			break
		}
		if arr[mid] > t { //Средний элемент маленький
			highBound = mid - 1
		}
		if arr[mid] < t { //Средний элемент большой
			lowBound = mid + 1
		}
	}
	return ans
}
