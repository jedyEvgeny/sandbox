//На вход подаётся число чисел и далее числа. Нужно из каждого входа удалить одну цифру, чтобы число оставалось максимальным.
// https://techpoint.ozon.ru/contest/173/task/329

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bufMaxSize := 256 * 1000 * 1000
	buf := make([]byte, 0)
	scanner.Buffer(buf, bufMaxSize)

	for count := 0; ; count++ {
		if !scanner.Scan() {
			break
		}
		if count == 0 {
			continue
		}
		line := scanner.Text()
		result := removeDigitToMaxNumber(line)

		if result != "" {
			fmt.Printf("%s\n", result)
		}
	}
}

func removeDigitToMaxNumber(line string) string {


	length := len(line)
	if length == 1 {
		return "0"
	}

	for i := 0; i < length; i++ {
		if i < length-1 && line[i] < line[i+1] {
			newNum := line[:i] + line[i+1:]
			return newNum
		}
	}
	return line[:length-1]
}
