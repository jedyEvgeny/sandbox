package main

import "fmt"

func main() {
	var (
		b     bool
		i     int
		str   string
		arr   [5]rune
		slice []byte
		ch    chan uint
		m     map[uint64]float64
		in    interface{ get() string } //можно без методов
		ptr   *int
		f     func()
		s     struct{ kv int } //можно без полей
	)
	fmt.Printf("Тип булевой: %T, значение булевой: %v\n", b, b)         //Вывод: Тип булевой: bool, значение булевой: false
	fmt.Printf("Тип числа: %T, значение числа: %v\n", i, i)             //Вывод: Тип числа: int, значение числа: 0
	fmt.Printf("Тип строки: %T, значение строки: %v\n", str, str)       //Вывод: Тип строки: string, значение строки:
	fmt.Printf("Тип массив 5: %T, значение массив 5: %v\n", arr, arr)   //Вывод: Тип массив 5: [5]int32, значение массив 5: [0 0 0 0 0]
	fmt.Printf("Тип срез: %T, значение срез: %v\n", slice, slice)       //Вывод: Тип срез: []uint8, значение срез: []
	fmt.Printf("Тип карта: %T, значение карта: %v\n", m, m)             //Вывод: Тип карта: map[uint64]float64, значение карта: map[]
	fmt.Printf("Тип структуры: %T, значение структуры: %v\n", s, s)     //Вывод: Тип структуры: struct { kv int }, значение структуры: {0}
	fmt.Printf("Тип канал: %T, значение канал: %v\n", ch, ch)           //Вывод: Тип канал: chan uint, значение канал: <nil>
	fmt.Printf("Тип интерфейса: %T, значение интерфейса: %v\n", in, in) //Вывод: Тип интерфейса: <nil>, значение интерфейса: <nil>
	fmt.Printf("Тип указателя: %T, значение указателя: %v\n", ptr, ptr) //Вывод: Тип указателя: *int, значение указателя: <nil>
	fmt.Printf("Тип функции: %T, значение функции: %v\n", f, f)         //Вывод: Тип функции: func(), значение функции: <nil>
}
