package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	mapN := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var ans int
		fmt.Scan(&ans)
		mapN[ans] = struct{}{}
	}
	sliceN := make([]int, 0)
	for idx := range mapN {
		sliceN = append(sliceN, idx)
	}
	sort.Ints(sliceN)
	for _, el := range sliceN {
		fmt.Println(el)
	}
}
