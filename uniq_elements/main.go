package main

import (
	"fmt"
	"log"
)

func main() {
	var lenSlice int
	_, err := fmt.Scan(&lenSlice)
	if err != nil {
		log.Fatal(err)
	}
	sliceNum := make([]int, lenSlice)
	mapUniq := make(map[int]uint)
	for i := 0; i < lenSlice; i++ {
		_, err := fmt.Scan(&sliceNum[i])
		if err != nil {
			log.Fatal(err)
		}
		mapUniq[sliceNum[i]]++
	}
	uniq := findUniq(mapUniq)
	fmt.Println(uniq)
}

func findUniq(m map[int]uint) uint {
	var countUniq uint
	for _, el := range m {
		if el < 2 {
			countUniq++
		}
	}
	return countUniq
}
