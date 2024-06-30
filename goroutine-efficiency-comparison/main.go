package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	count = 50
)

func calculatePi() float64 {
	insideCircle := 0
	totalPoints := 0

	for i := 0; i < 1000000; i++ {
		x := randFloat()
		y := randFloat()

		if math.Pow(x, 2)+math.Pow(y, 2) <= 1.0 {
			insideCircle++
		}
		totalPoints++
	}

	return 4.0 * float64(insideCircle) / float64(totalPoints)
}

func randFloat() float64 {
	return rand.Float64()
}

func main() {
	runtime.GOMAXPROCS(16)

	var piWithoutGoroutines float64
	progressBar := "####################"
	fmt.Println("Прогресс решения:")
	startTime := time.Now()
	for i := 0; i < count; i++ {
		fmt.Printf("\r[%-20s] %3d%%", progressBar[:i*20/count+1], (i+1)*100/count)
		piWithoutGoroutines = calculatePi()
	}
	endTime := time.Since(startTime)
	fmt.Printf(".\nРезультат без использования горутин: %f за время %v.\n", piWithoutGoroutines, endTime)

	var wg sync.WaitGroup
	wg.Add(count)
	var piWithGoroutines float64
	startTime = time.Now()
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			piWithGoroutines = calculatePi()
		}()
	}
	wg.Wait()
	endTimeGoroutines := time.Since(startTime)
	fmt.Printf("Результат с использованием горутин: %f за время %v.\n", piWithGoroutines, endTimeGoroutines)

	efficiency := float64(endTime.Nanoseconds()) / float64(endTimeGoroutines.Nanoseconds())
	timeDifference := endTime.Seconds() - endTimeGoroutines.Seconds()
	fmt.Println("------------------")
	fmt.Printf("Код с горутинами эффективнее в %.2f раз и быстрее на %.3f секунды.\n", efficiency, timeDifference)
}
