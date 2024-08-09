// Определяем сколько время в другом городе
// Была задача: писать ежедневный отчёт о проделанной работе,
// в т.ч. с указанием времени работы по мск.
// Чтобы каждый раз не считать вручную, написал простенькую программу
package main

import (
	"fmt"
	"log"
)

type timeZone struct {
	myZone          string
	destinationZone string
	timeDifference  uint8
}

func main() {
	fmt.Println("---Сервис сравнения времени---")
	var t = timeZone{}
	fmt.Println("Введите название вашего места:")
	fmt.Scan(&t.myZone)
	fmt.Println("Введите название второго места:")
	fmt.Scan(&t.destinationZone)
	fmt.Println("Введите разницу часовых поясов:")
	fmt.Scan(&t.timeDifference)
	fmt.Println("----------")
	var ans int
	for {
		fmt.Println("Введите ваше время в часах:")
		_, err := fmt.Scan(&ans)
		if err != nil {
			log.Println("Введено не время в часах")
		}
		fmt.Printf("%s: %d час.\n", t.destinationZone, ans-int(t.timeDifference))
	}
