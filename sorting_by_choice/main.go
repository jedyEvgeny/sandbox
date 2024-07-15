package main

import "fmt"

func main() {
	slice := []int{2, -5, 6, -15, 14, 225, 64}
	sliceSec := make([]int, len(slice))
	copy(sliceSec, slice)
	newSlice := sortSlice(slice)
	fmt.Println(newSlice)
	selectionSort(sliceSec)
	fmt.Println(sliceSec)
}

func sortSlice(arr []int) []int {
	newSlice := make([]int, len(arr))
	sizeSlice := len(arr) //Длина слайса arr уменьшается в цикле
	for i := 0; i < sizeSlice; i++ {
		smallestIdxOfElem := findSmallestIdx(arr) //Ищем наименьший элемент
		newSlice[i] = arr[smallestIdxOfElem]      //Добавляем наименьший элемент в новый слайс
		arr = append(arr[:smallestIdxOfElem], arr[smallestIdxOfElem+1:]...)
	}
	return newSlice
}

func findSmallestIdx(arr []int) int {
	smallestEl := arr[0] //Храним наименьший элемент
	var smallestIdx int  //Храним индекс наименьшего элемента
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallestEl {
			smallestEl = arr[i]
			smallestIdx = i
		}
	}
	return smallestIdx
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] //Синтаксический сахар
	}
}
