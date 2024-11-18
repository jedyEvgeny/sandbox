//Тренерую односвязные списки
package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type singlyLinkedList struct {
	head *node
}

func main() {
	l := singlyLinkedList{}
	list := &l
	for i := 1; i < 4; i++ {
		list.add(i)
	}
	list.print()

	for i := 0; i < 4; i++ {
		node, ok := list.findValNodePerIdx(i)
		if ok {
			fmt.Printf("Значение %d-го элемента очереди равено: %d\n",
				i, node.value)
		}
		if !ok {
			fmt.Printf("%d-й элемент в очереди отсутствует\n", i)
		}
	}
	list.reverse()
	list.print()

}

func (l *singlyLinkedList) add(value int) {
	n := node{value: value}
	node := &n
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
	strSep := strings.Repeat("-", 30)
	fmt.Println(strSep)
	currNode := l.head
	for i := 0; currNode != nil; i++ {
		fmt.Printf("Значение %d-го элемента списка равен: %d\n",
			i, currNode.value)
		currNode = currNode.next
	}
	fmt.Println(strSep)
}

func (l *singlyLinkedList) findValNodePerIdx(idx int) (*node, bool) {
	currNode := l.head
	for i := 0; currNode != nil; i++ {
		if i == idx {
			return currNode, true
		}
	}
	return nil, false
}

func (l *singlyLinkedList) rev() {
	var prev *node
	currNode := l.head
	for currNode != nil {
		currNode.next, prev, currNode = prev, currNode, currNode.next
	}
}

func (l *singlyLinkedList) reverse() {
	var prev *node
	currNode := l.head

	for currNode != nil {
		next := currNode.next //Сохраняем следующий узел
		currNode.next = prev  //Присваиваем в текущем узле полю сл. узла последний узел
		prev = currNode       //Присваиваем переменной текущий узел
		currNode = next       //Переходим к следующему узлу
	}
	l.head = prev
}
