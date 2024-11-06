//Тренируюсь работать с односвязным списком

package main

import "fmt"

// Node определяет узел односвязного списка
type Node struct {
	Value int
	Next  *Node
}

// SinglyLinkedList предоставляет односвязный список
type SinglyLinkedList struct {
	Head *Node
}

func main() {
	list := &SinglyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	fmt.Print("Содержимое списка: ")
	list.Print()

	if node, found := list.Get(1); found {
		fmt.Println("Элемент на индексе 1:", node.Value)
	} else {
		fmt.Println("Элемент не найден на индексе 1")
	}

	if node, found := list.Get(5); found {
		fmt.Println("Элемент на индексе 5:", node.Value)
	} else {
		fmt.Println("Элемент не найден на индексе 5")
	}
}

// Add добавляет новый элемент в конец списка
func (list *SinglyLinkedList) Add(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	now := list.Head
	for now.Next != nil {
		now = now.Next
	}
	now.Next = newNode
}

// Print печатает элементы односвязного списка
func (list *SinglyLinkedList) Print() {
	now := list.Head
	for now != nil {
		fmt.Print(now.Value, " ")
		now = now.Next
	}
	fmt.Println()
}

func (list *SinglyLinkedList) Get(idx int) (*Node, bool) {
	now := list.Head
	for i := 0; now != nil; i++ {
		if i == idx {
			return now, true
		}
		now = now.Next
	}
	return nil, false
}
