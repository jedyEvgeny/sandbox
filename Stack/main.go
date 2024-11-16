//Реализую стек на го
package main

import "fmt"

type stack struct {
	values []int
}

func main() {
	stack := &stack{}
	for i := 0; i < 10000000; i++ {
		stack.push(i)
	}
	fmt.Println("Длина стека:", len(stack.values))
	topElem, _ := stack.pop()
	fmt.Println("Верхний элемент стека:", topElem)
	fmt.Println("Длина стека после удаления одного элемента:",
		len(stack.values))
}

func (s *stack) push(value int) {
	s.values = append(s.values, value)
}

func (s *stack) pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	index := len(s.values) - 1
	element := s.values[index]
	s.values = s.values[:index]
	return element, true
}
