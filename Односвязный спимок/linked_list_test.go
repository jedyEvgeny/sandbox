package main

import (
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := &SinglyLinkedList{}

	// Тест на пустой список
	if _, found := list.Get(0); found {
		t.Error("Expected not found for empty list")
	}

	// Тест на добавление первого элемента
	list.Add(1)
	if node, found := list.Get(0); !found || node.Value != 1 {
		t.Errorf("Expected 1 as first element, got %v", node)
	}

	// Тест на добавление нескольких элементов
	list.Add(2)
	list.Add(3)
	if node, found := list.Get(1); !found || node.Value != 2 {
		t.Errorf("Expected 2 at index 1, got %v", node)
	}
	if node, found := list.Get(2); !found || node.Value != 3 {
		t.Errorf("Expected 3 at index 2, got %v", node)
	}

	// Тест на извлечение по несуществующему индексу
	if _, found := list.Get(3); found {
		t.Error("Expected not found for index 3")
	}

	// Тест на последовательность значений
	expectedValues := []int{1, 2, 3}
	for i, expected := range expectedValues {
		if node, found := list.Get(i); !found || node.Value != expected {
			t.Errorf("Expected %d at index %d, got %v", expected, i, node)
		}
	}

	// (Дополнительно) Тест на удаление - потребуется реализация функции удаления
	// list.Remove(2)
	// if _, found := list.Get(1); found {
	//     t.Error("Expected not found for index 1 after removal")
	// }
}
