package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 5, 19}
	uniq := findUnic(arr)
	fmt.Println(uniq)
}

func findUnic(arr []int) int {
	mapUnic := make(map[int]struct{})
	var count int
	for _, val := range arr {
		_, ok := mapUnic[val]
		if !ok {
			mapUnic[val] = struct{}{}
			arr[count] = val
			count++
		}
	}
	return len(mapUnic)
}
