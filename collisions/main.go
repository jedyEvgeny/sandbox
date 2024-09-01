//Программа для искусственного поиска коллизий, то есть одинакового хеша для разных входных данных
//На 42м миллионе через несколько минут работы, IDE (с терминалом) зависла, но весь компьютер продолжал работать. 
// Предположу, что можно было оставить работать, коллизия бы нашлась. Либо нужно искать в строках, а не числах 

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strconv"
)

func main() {
	h := sha1.New()
	_, err := io.WriteString(h, "100")
	hash := fmt.Sprintf("%x", h.Sum(nil))
	if err != nil {
		log.Println(err)
		return
	}

	hashes := make(map[string]bool)
	hashes[hash] = false

	count := 101
	var newHash string

	for {
		count++
		countStr := strconv.Itoa(count)
		h := sha1.New()
		_, err = io.WriteString(h, countStr)
		if err != nil {
			log.Fatal(err)
		}
		newHash = fmt.Sprintf("%x", h.Sum(nil))
		value, ok := hashes[newHash]
		if ok {
			hashes[newHash] = true
			break
		}
		fmt.Println(count, value, newHash)
	}
	fmt.Printf("Исходный хеш числа %d: %s\n", 100, hash)
	fmt.Printf("Хеш-дубликат числа %d: %s\n", count, hash)
}

//5916024
