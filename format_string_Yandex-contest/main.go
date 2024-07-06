package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// input := "b,,,,,,,,,,,,dc,,,,,"
	// input := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,yandex"
	// input := "once upon a time, in a land far far away lived a princess , whose beauty was yet unmatched,  dd fdffdsdf,"

	sliceWords := strings.FieldsFunc(input, func(r rune) bool {
		return !('a' <= r && r <= 'z')
	})

	var bestWord int
	for _, word := range sliceWords {
		if len(word) > bestWord {
			bestWord = len(word)
		}
	}

	words := strings.SplitAfter(input, ",")
	str := strings.Join(words, " ")
	sliceWithoutSpaces := strings.Fields(str)
	maxLength := bestWord * 3

	var currentLine string

	for _, word := range sliceWithoutSpaces {

		if len(currentLine)+len(word) <= maxLength && (word == "," || currentLine == "") {
			currentLine += word
			continue
		}

		if len(currentLine)+len(word)+1 <= maxLength {
			currentLine += " " + word
			continue
		}

		fmt.Println(currentLine)
		currentLine = word
	}

	if len(currentLine) > 0 {
		fmt.Println(currentLine)
	}
}
