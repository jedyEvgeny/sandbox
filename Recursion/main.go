//Опыты разработки алгоритмов на рекурсию

package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{5, 88, 194, 329, 531, 624}
	fmt.Println("Суммирование:", sumReq(arr))
	fmt.Println("Элементов среза:", countElemReq(arr))
	fmt.Println("Максимальное число среза:", maxReq(arr))
	fmt.Println("Бинарный поиск числа:", binSearchReq(arr, 531, 0, len(arr)-1))
}

//sumReq суммирует элементы среза
func sumReq(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := sumReq(arr[1:])
	log.Println(sum)
	return arr[0] + sum
}

/*
1 2 3 - в стеке общая функция
2 3   - в стеке общая функция и функция на 2-3
3	  - в стеке общая функция, функция на 2-3 и функция на 3
[]int - в стеке общая функция, функция на 2-3, функция на 3 и функция на нулевой срез
3 	  - наступил базовый случай, извлекаем из стека первое значение
+3	  - суммируем значение базового случая с нулевым элементом текущего массива
2+3	  - извлекаем следующий слой стека
1+2+3 - извлекаем следующий слой стека и возвращаем сумму
*/

func countElemReq(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + countElemReq(arr[1:])
}

//maxReq ищет максимальное значение элемента среза
func maxReq(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	subMax := maxReq(arr[1:])
	log.Println(subMax)
	if arr[0] > subMax {
		return arr[0]
	}
	return subMax
}

//binSearchReq выполняет бинарный поиск
func binSearchReq(slice []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if slice[mid] == target {
		return mid
	}

	if slice[mid] < target {
		return binSearchReq(slice, target, mid+1, right)
	}

	return binSearchReq(slice, target, left, mid-1)
}
