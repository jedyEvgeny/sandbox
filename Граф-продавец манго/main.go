/*
Программа проверяет, есть ли среди ваших знакомых или знакомых ваших знакомых нужная персона
т.е. - есть ли связь
Пример программы на графы
*/
package main

import "fmt"

func main() {
	graph := make(map[string][]string)

	graph["Вы"] = []string{"Человек", "Покемон", "Горыныч"}

	graph["Человек"] = []string{"Лиса", "Медведь", "Тигр", "Вы"}
	graph["Покемон"] = []string{"Дигимон", "Эш", "Эшли", "Вы"}
	graph["Горыныч"] = []string{"Яга", "Василиса", "Иван", "Вы"}

	graph["Лиса"] = []string{"Орк", "Гоблин", "Человек"}
	graph["Медведь"] = []string{"Пчела", "Медведица", "Человек"}
	graph["Тигр"] = []string{"Леопард", "Человек"}

	graph["Дигимон"] = []string{"Покемон"}
	graph["Эш"] = []string{"Покемон"}
	graph["Эшли"] = []string{"Покемон"}

	graph["Яга"] = []string{"Горыныч"}
	graph["Василиса"] = []string{"Горыныч"}
	graph["Иван"] = []string{"Горыныч"}

	graph["Орк"] = []string{"Лиса"}
	graph["Гоблин"] = []string{"Лиса"}
	graph["Пчела"] = []string{"Медведь"}
	graph["Медведица"] = []string{"Медведь"}
	graph["Леопард"] = []string{"Тигр"}

	target := "Василиса"

	search("Вы", target, graph)
}

func search(start, target string, graph map[string][]string) {
	searchQueue := []string{start}
	var chekedPersons []string
	for len(searchQueue) > 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]
		if isInChekedPersons(person, chekedPersons) {
			continue
		}
		ok := isTarget(target, person)
		if ok {
			fmt.Printf("Нашли связь с %s\n", target)
			return
		}
		searchQueue = append(searchQueue, graph[person]...)
		chekedPersons = append(chekedPersons, person)
	}
	fmt.Println("Связь не найдена")
}

func isInChekedPersons(person string, chekedPersons []string) bool {
	for _, el := range chekedPersons {
		if el == person {
			return true
		}
	}
	return false
}

func isTarget(target, person string) bool {
	return target == person
}
