package main

import (
	"fmt"
	"math"
)

type city struct {
	coordinateX, coordinateY, id int
}

func main() {
	var amountCities, x, y, from, to int
	var distance float64
	fmt.Scan(&amountCities)
	cities := make([]city, amountCities)
	for i := 0; i < amountCities; i++ {
		fmt.Scan(&x, &y)
		cities[i] = city{
			coordinateX: x,
			coordinateY: y,
			id:          i + 1,
		}
	}
	fmt.Scan(&distance)
	fmt.Scan(&from, &to)
	roads := minRoads(cities, distance, from, to)
	fmt.Println(roads)
}

func minRoads(cities []city, maxDistance float64, start, end int) int {
	queue := make([]int, 0)
	visited := make(map[int]bool)

	queue = append(queue, start)
	visited[start] = true

	roads := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == end {
				return roads
			}
			neighbors := getReachableCities(cities, maxDistance, current)
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		roads++
	}
	return -1
}

func getReachableCities(cities []city, maxDistance float64, start int) []int {
	reachableCities := []int{start}
	for i, city := range cities {
		if i != start-1 {
			if distance(cities[start-1], city) <= maxDistance {
				reachableCities = append(reachableCities, i+1)
			}
		}
	}
	return reachableCities
}

func distance(c1, c2 city) float64 {
	return math.Abs(float64(c1.coordinateX-c2.coordinateX)) + math.Abs(float64(c1.coordinateY-c2.coordinateY))
}
