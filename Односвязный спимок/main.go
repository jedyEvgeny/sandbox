package main

import "fmt"

type node struct {
	value int
	next  *node
}

type singlyLinkedList struct {
	head *node
}

func main() {
	l := &singlyLinkedList{}
	for i := 100; i < 104; i++ {
		l.add(i)
	}
	l.print()

	idx := 2
	val, ok := l.findNode(idx)
	printVal(idx, val, ok)

	l.reverse()
	l.print()

	l.delete(idx)
	l.print()
}

func (l *singlyLinkedList) add(val int) {
	node := &node{value: val}
	if l.head == nil {
		l.head = node
		return
	}

	currNode := l.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = node
}

func (l *singlyLinkedList) print() {
	fmt.Printf("Элементы списка:")
	currNode := l.head
	for currNode != nil {
		fmt.Printf("%d->", currNode.value)
		currNode = currNode.next
	}
	fmt.Printf("nil\n\n")
}

func (l *singlyLinkedList) findNode(idx int) (int, bool) {
	currNode := l.head
	for i := 0; currNode != nil && i <= idx; i++ {
		if i == idx {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

func printVal(idx, val int, ok bool) {
	if ok {
		fmt.Printf("Для %d-го элемента значение равно: %d\n\n",
			idx, val)
	}
	if !ok {
		fmt.Printf("%d-й элемент в списке отсутствует\n\n", idx)
	}
}

func (l *singlyLinkedList) reverse() {
	var prev *node
	currNode := l.head
	for currNode != nil {
		next := currNode.next
		currNode.next = prev
		prev = currNode
		currNode = next
	}
	l.head = prev
}

func (l *singlyLinkedList) delete(idx int) {
	if l.head == nil {
		fmt.Println("Список пуст, нечего удалять.")
		return
	}

	if idx == 0 {
		l.head = l.head.next
		return
	}

	prevNode := l.head
	currNode := l.head.next
	for i := 1; currNode != nil && i < idx; i++ {
		prevNode = currNode
		currNode = currNode.next
	}

	if currNode == nil {
		fmt.Printf("Индекс %d выходит за пределы списка.\n", idx)
		return
	}

	prevNode.next = currNode.next
}
