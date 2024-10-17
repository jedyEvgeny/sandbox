//Задача с Яндекса.
//Суть: из файла поступают данные - количество людей и данные о людях в формате: Volog,Igor,Danilovich,3,14,1962
//Нужно зашифровать каждую строку по особой схеме и вывести в терминал через зпт, в формате: 80D, 48F, 184, 710, 64F
//Задание уже не ищу, схему шифровки можно помотреть по коду
//Из интересного: если в цикле считать в массив данные, а затем циклом перебрать каждую строку - тут бенчмарк будет примерно такой же, если всё сделать за один цикл. Странности.
//Запускать программу так: cat input.txt | go run main.go
//Запустить бенчмарк так: cat in.txt | go test -bench .

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type people struct {
	sername     string
	name        string
	patronymic  string
	birthDay    string
	birthMounth string
	birthYear   int
}

const mulSumDay = 64

func main() {
	run()
}

// func run() {
// 	var lengthSlice int
// 	_, _ = fmt.Scan(&lengthSlice)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	peoples := make([]people, lengthSlice)
// 	//fmt.Println(len(peoples))
// 	var count int
// 	for {
// 		if !scanner.Scan() {
// 			break
// 		}
// 		data := scanner.Text()
// 		i := parseData(data)
// 		peoples[count] = i
// 		count++
// 	}
// 	//fmt.Println("-------")
// 	for idx, person := range peoples {
// 		uniqSyb := countUniqSym(person)
// 		//fmt.Println("Уникальных символов в ФИО:", uniqSyb)

// 		numDayAndMounth := sumDays(person)
// 		//fmt.Println("Сумма цифр дня и месяца:", numDayAndMounth)

// 		idFirstLetter := idFirstLetter(person.sername)
// 		//fmt.Println("Порядок буквы:", idFirstLetter)

// 		totalSumStr := sumNums(uniqSyb, numDayAndMounth, idFirstLetter)
// 		//fmt.Println("16х:", totalSumStr)

// 		lastThree := lastThree(totalSumStr)
// 		//fmt.Println("ID:", lastThree)
// 		//fmt.Println("---------")
// 		if idx == lengthSlice-1 {
// 			fmt.Print(lastThree)
// 			break
// 		}
// 		fmt.Print(lastThree + ", ")
// 	}
// }

func run() {
	var lengthSlice int
	_, _ = fmt.Scan(&lengthSlice)
	scanner := bufio.NewScanner(os.Stdin)
	// peoples := make([]people, lengthSlice)
	//fmt.Println(len(peoples))
	var count int
	for {
		if !scanner.Scan() {
			break
		}
		data := scanner.Text()
		person := parseData(data)
		// peoples[count] = i

		uniqSyb := countUniqSym(person)
		//fmt.Println("Уникальных символов в ФИО:", uniqSyb)

		numDayAndMounth := sumDays(person)
		//fmt.Println("Сумма цифр дня и месяца:", numDayAndMounth)

		idFirstLetter := idFirstLetter(person.sername)
		//fmt.Println("Порядок буквы:", idFirstLetter)

		totalSumStr := sumNums(uniqSyb, numDayAndMounth, idFirstLetter)
		//fmt.Println("16х:", totalSumStr)

		lastThree := lastThree(totalSumStr)
		//fmt.Println("ID:", lastThree)
		//fmt.Println("---------")
		if count == lengthSlice-1 {
			fmt.Print(lastThree)
			break
		}
		fmt.Print(lastThree + ", ")

		count++
	}
}

func parseData(data string) people {
	slice := strings.Split(data, ",")
	return people{
		sername:     slice[0],
		name:        slice[1],
		patronymic:  slice[2],
		birthDay:    slice[3],
		birthMounth: slice[4],
		birthYear:   parseDataBirthday(slice[5]),
	}
}

func parseDataBirthday(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func countUniqSym(p people) int {
	m := make(map[rune]struct{})
	str := p.sername + p.name + p.patronymic
	//fmt.Println(str)
	for _, sym := range str {
		m[sym] = struct{}{}
	}
	return len(m)
}

func sumDays(p people) int {
	nums := make([]int, 0, 4)
	nums = parseDigits(p.birthDay, nums)
	nums = parseDigits(p.birthMounth, nums)
	var sum int
	//fmt.Println("Массив:", nums)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func parseDigits(str string, nums []int) []int {
	for _, el := range str {
		nums = append(nums, int(el-'0'))
	}
	return nums
}

func idFirstLetter(str string) int {
	letter := rune(str[0])
	return int(letter - '@')
}

func sumNums(n1, n2, n3 int) string {
	sum := n1 + n2*mulSumDay + n3*256
	//fmt.Println("Сумма по-формуле:", sum)
	return fmt.Sprintf("%X", sum)
}

func lastThree(str string) string {
	length := len(str)
	if length == 3 {
		return str
	}
	switch length {
	case 0:
		str = "000"
	case 1:
		str = "00" + str
	case 2:
		str = "0" + str
	default:
		str = str[length-3:]
	}
	return str
}
