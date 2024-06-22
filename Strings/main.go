// Время выполнения эффективной работы: 62.923µs
// Время выполнения неэффективной работы: 709.326µs
package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	// Пример эффективной работы со строками через буфер
	startTimeEfficient := time.Now()

	var buf bytes.Buffer

	for i := 0; i < 1000; i++ {
		buf.WriteString(string(i))
	}
	resultEfficient := buf.String()

	elapsedTimeEfficient := time.Since(startTimeEfficient)
	fmt.Println("Время выполнения эффективной работы:", elapsedTimeEfficient)

	// Пример неэффективной работы со строками через оператор "+"
	startTimeInefficient := time.Now()

	var resultInefficient string
	for i := 0; i < 1000; i++ {
		resultInefficient += string(i)
	}

	elapsedTimeInefficient := time.Since(startTimeInefficient)
	fmt.Println("Время выполнения неэффективной работы:", elapsedTimeInefficient)
	fmt.Println(resultEfficient)
	fmt.Println()
	fmt.Println(resultInefficient)
}
