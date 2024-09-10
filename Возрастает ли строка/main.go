/*Задача с CodeRun уровня easy
https://coderun.yandex.ru/selections/quickstart/problems/list-growing/description?currentPage=1&pageSize=20&search=
Дан список. Определите, является ли он монотонно возрастающим (то есть верно ли, что каждый элемент этого списка больше предыдущего).
Выведите YES, если массив монотонно возрастает и NO в противном случае.

Вызовом для меня было - не читать весь поток из стандартного ввода, 
т.к. он теоретически может занимать гигабайты памяти, при этом иметь в начале потока условие прекращения чтения,
а поэлементно считывать по пробелу.

Проблема заключалась в том, что нет понимания - сколько знаков нужно считывать.
К этому времени я уже познакомился со стандартным NewScanner. В ходе решения познакомился с NewReader, где с помощью разделителя по пробелу
считывал поэлементно - но NewReader не считывал последний элемент потока, т.к. за ним не было пробела.

Решил задачу с помощью того же NewScanner, поторатив на это часа три
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var lastNum int

	for i := 0; ; i++ {
		if !scanner.Scan() {
			break
		}
		numStr := scanner.Text()
		numStr = strings.TrimSpace(numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 || num > lastNum {
			lastNum = num
			continue
		}
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
