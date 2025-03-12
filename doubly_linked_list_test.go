package main

import (
	"testing"
)

func TestLength(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.Length() != 0 {
		t.Errorf("Очікувана довжина 0, отримано %d", list.Length())
	}

	list.Append('a')
	list.Append('b')
	if list.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", list.Length())
	}
}

func TestAppend(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')

	if list.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", list.Length())
	}

	value, _ := list.Get(1)
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}
}

func TestInsert(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('c')

	err := list.Insert('b', 1)
	if err != nil {
		t.Errorf("Неочікувана помилка: %v", err)
	}

	value, _ := list.Get(1)
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}

	err = list.Insert('d', 10) // Неправильний індекс
	if err == nil {
		t.Error("Очікувана помилка, отримано nil")
	}
}

func TestDelete(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')
	list.Append('c')

	value, err := list.Delete(1)
	if err != nil {
		t.Errorf("Неочікувана помилка: %v", err)
	}
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}

	if list.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", list.Length())
	}

	_, err = list.Delete(10) // Неправильний індекс
	if err == nil {
		t.Error("Очікувана помилка, отримано nil")
	}
}

func TestDeleteAll(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')
	list.Append('a')
	list.Append('c')

	list.DeleteAll('a')
	if list.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", list.Length())
	}

	value, _ := list.Get(0)
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}
}

func TestGet(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')

	value, err := list.Get(1)
	if err != nil {
		t.Errorf("Неочікувана помилка: %v", err)
	}
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}

	_, err = list.Get(10) // Неправильний індекс
	if err == nil {
		t.Error("Очікувана помилка, отримано nil")
	}
}

func TestClone(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')

	clone := list.Clone()
	if clone.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", clone.Length())
	}

	value, _ := clone.Get(1)
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}

	list.Append('c') // Зміна оригінального списку не повинна впливати на клон
	if clone.Length() != 2 {
		t.Errorf("Очікувана довжина 2, отримано %d", clone.Length())
	}
}

func TestReverse(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')
	list.Append('c')

	list.Reverse()

	value, _ := list.Get(0)
	if value != 'c' {
		t.Errorf("Очікуваний елемент 'c', отримано %c", value)
	}

	value, _ = list.Get(1)
	if value != 'b' {
		t.Errorf("Очікуваний елемент 'b', отримано %c", value)
	}

	value = list.head.value
	if value != 'c' {
		t.Errorf("Очікуваний елемент 'c', отримано %c", value)
	}

	value = list.tail.value
	if value != 'a' {
		t.Errorf("Очікуваний елемент 'a', отримано %c", value)
	}
}

func TestFindFirst(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')
	list.Append('a')

	index := list.FindFirst('a')
	if index != 0 {
		t.Errorf("Очікуваний індекс 0, отримано %d", index)
	}

	index = list.FindFirst('c')
	if index != -1 {
		t.Errorf("Очікуваний індекс -1, отримано %d", index)
	}
}

func TestFindLast(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')
	list.Append('a')

	index := list.FindLast('a')
	if index != 2 {
		t.Errorf("Очікуваний індекс 2, отримано %d", index)
	}

	index = list.FindLast('t')
	if index != -1 {
		t.Errorf("Очікуваний індекс -1, отримано %d", index)
	}
}

func TestClear(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Append('a')
	list.Append('b')

	list.Clear()
	if list.Length() != 0 {
		t.Errorf("Очікувана довжина 0, отримано %d", list.Length())
	}
}

func TestExtend(t *testing.T) {
	list1 := NewDoublyLinkedList()
	list1.Append('a')
	list1.Append('b')

	list2 := NewDoublyLinkedList()
	list2.Append('c')
	list2.Append('d')

	list1.Extend(list2)
	if list1.Length() != 4 {
		t.Errorf("Очікувана довжина 4, отримано %d", list1.Length())
	}

	value, _ := list1.Get(2)
	if value != 'c' {
		t.Errorf("Очікуваний елемент 'c', отримано %c", value)
	}
}
