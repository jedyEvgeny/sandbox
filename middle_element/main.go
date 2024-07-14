package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 3)
	for i := 0; i < 3; i++ {
		var ans int
		fmt.Scan(&ans)
		arr[i] = ans
	}
	sort.Ints(arr)
	fmt.Println(arr[1])
}
