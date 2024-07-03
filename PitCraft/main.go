package main

import (
	"fmt"
)

func calculateWaterTrapped(blocks []int) int {
	length := len(blocks)
	if length <= 2 {
		return 0
	}

	leftMax := make([]int, length)
	rightMax := make([]int, length)

	leftMax[0] = blocks[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], blocks[i])
	}

	rightMax[length-1] = blocks[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], blocks[i])
	}

	totalWater := 0
	for i := 1; i < length-1; i++ {
		minHeight := min(leftMax[i], rightMax[i])
		if minHeight > blocks[i] {
			totalWater += minHeight - blocks[i]
		}
	}

	return totalWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printBlocks(blocks []int) {
	maxHeight := getMax(blocks)
	for h := maxHeight; h > 0; h-- {
		for _, block := range blocks {
			if block >= h {
				fmt.Printf("▓▓")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}

func getMax(blocks []int) int {
	maxHeight := 0
	for _, block := range blocks {
		if block > maxHeight {
			maxHeight = block
		}
	}
	return maxHeight
}

func main() {
	blocks := []int{3, 1, 4, 2, 3, 1}

	printBlocks(blocks)

	trappedWater := calculateWaterTrapped(blocks)
	fmt.Printf("Количество задержавшейся воды: %d\n", trappedWater)
}
