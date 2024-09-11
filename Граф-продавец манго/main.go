/*
Программа ищет продавца манго среди ваших друзей и друзей ваших друзей
Условие для продавца манго - фамилия заканчисается на m
Пример программы на графы
*/

package main

var graph = make(map[string][]string)

func main() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	search("you")
}

func search(name string) bool {
	var searchQueue []string //Очередь для поиска
	searchQueue = append(searchQueue, graph[name]...)

	var chekedPersons []string //уже проверенные люди
	var person string

	for len(searchQueue) != 0 { //Пока очередь не пуста
		person = searchQueue[0]       //Проверяем первую персону в очереди
		searchQueue = searchQueue[1:] //Выводим из очереди текущую персону
		if isInChekedPerson(person, chekedPersons) {
			continue
		}
		if isSeller(person) {
			println(person + " is mango seller!")
			return true
		}

		//расширяем очередь новыми связями
		searchQueue = append(searchQueue, graph[person]...)

		//добавляем проверенную персону в список проверенных
		chekedPersons = append(chekedPersons, person)
	}
	return false
}

func isInChekedPerson(person string, searched []string) bool {
	for _, n := range searched {
		if n == person {
			return true
		}
	}
	return false
}

func isSeller(name string) bool {
	return name[len(name)-1] == 'm'
}
