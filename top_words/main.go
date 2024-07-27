// Имеется несколько файлов. Нужно прочитать файлы и вывести топ-10 слов
// Запускаем с аргументом, примерно так: go run main.go /.short_files
// Символ ./ используется в bash-языке как символ относительного пути к текущему каталогу
// ЗЫ - нужно создать файлы с содержимым
// Лучшее время набора кода 14 минут
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	pathDirFiles := "/files"
	if len(os.Args) == 2 {
		pathDirFiles = os.Args[1]
	}
	filesList, err := os.ReadDir(pathDirFiles)
	if err != nil {
		log.Println("не удаётся найти каталог:", err)
		return
	}
	var allWordsSlice []string
	for _, fileEntry := range filesList {
		if fileEntry.IsDir() {
			continue
		}
		fullFilePath := filepath.Join(pathDirFiles, fileEntry.Name())
		contentFile, err := os.ReadFile(fullFilePath)
		if err != nil {
			log.Printf("не удалось прочитать файл %s: %s", fileEntry.Name(), err)
			continue
		}
		words := strings.Fields(string(contentFile))
		allWordsSlice = append(allWordsSlice, words...)
	}
	allWordsMap := make(map[string]int)
	for _, value := range allWordsSlice {
		allWordsMap[value]++
	}
	type topWords struct {
		word      string
		frequency int
	}
	var frequencyWordsSlice []topWords
	for key, value := range allWordsMap {
		frequencyWordsSlice = append(frequencyWordsSlice, topWords{key, value})
	}
	sort.Slice(frequencyWordsSlice, func(i, j int) bool {
		return frequencyWordsSlice[i].frequency > frequencyWordsSlice[j].frequency
	})
	for i := 0; i < 10 && i < len(frequencyWordsSlice); i++ {
		fmt.Printf("слово (%s)\t встречается %d раз\n", frequencyWordsSlice[i].word, frequencyWordsSlice[i].frequency)
	}
}
