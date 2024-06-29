package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	arrStr := strings.Fields(string(content))
	// arrStr := strings.Split(string(content), " ") //Некорректно отрабатывает в контесте
	numFirst, _ := strconv.Atoi(arrStr[0])
	numSecond, _ := strconv.Atoi(arrStr[1])

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	sum := numFirst + numSecond
	sumStr := strconv.Itoa(sum)
	_, err = outputFile.WriteString(sumStr)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}
