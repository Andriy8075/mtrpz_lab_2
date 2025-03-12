package main

import (
	"errors"
	"fmt"
)

type Character rune

type Node struct {
	value Character
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func NewList() *List {
	return &List{}
}

func (list *List) Length() int {
	return list.size
}

func (list *List) Append(element Character) {
	newNode := &Node{value: element, prev: list.tail}

	if list.tail == nil {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}

	list.tail = newNode
	list.size++
}

func (list *List) Insert(element Character, index int) error {
	if index < 0 || index > list.size {
		return errors.New("індекс поза межами діапазону")
	}

	if index == list.size {
		list.Append(element)
		return nil
	}

	newNode := &Node{value: element}
	if index == 0 {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}

	list.size++
	return nil
}

func (list *List) Delete(index int) (Character, error) {
	if index < 0 || index >= list.size {
		return 0, errors.New("індекс поза межами діапазону")
	}

	var value Character
	if list.size == 1 {
		value = list.head.value
		list.head = nil
		list.tail = nil
	} else if index == 0 {
		value = list.head.value
		list.head = list.head.next
		list.head.prev = nil
	} else if index == list.size-1 {
		value = list.tail.value
		list.tail = list.tail.prev
		list.tail.next = nil
	} else {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		value = current.value
		current.prev.next = current.next
		current.next.prev = current.prev
	}

	list.size--
	return value, nil
}

func (list *List) DeleteAll(element Character) {
	current := list.head
	for current != nil {
		if current.value == element {
			if current.prev == nil {
				list.head = current.next
				if list.head != nil {
					list.head.prev = nil
				}
			} else {
				current.prev.next = current.next
			}

			if current.next == nil {
				list.tail = current.prev
				if list.tail != nil {
					list.tail.next = nil
				}
			} else {
				current.next.prev = current.prev
			}

			list.size--
		}
		current = current.next
	}
}

func (list *List) Get(index int) (Character, error) {
	if index < 0 || index >= list.size {
		return 0, errors.New("індекс поза межами діапазону")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (list *List) Clone() *List {
	newList := NewList()
	current := list.head
	for current != nil {
		newList.Append(current.value)
		current = current.next
	}
	return newList
}

func (list *List) Reverse() {
	current := list.head
	list.tail = current
	var temp *Node

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		list.head = temp.prev

	}
}

func (list *List) FindFirst(element Character) int {
	current := list.head
	index := 0

	for current != nil {
		if current.value == element {
			return index
		}
		current = current.next
		index++
	}

	return -1
}

func (list *List) FindLast(element Character) int {
	current := list.tail
	index := list.size - 1
	for current != nil {
		if current.value == element {
			return index
		}
		current = current.prev
		index--
	}

	return -1
}

func (list *List) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *List) Extend(other *List) {
	current := other.head
	for current != nil {
		list.Append(current.value)
		current = current.next
	}
}

func (list *List) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%c ", current.value)
		current = current.next
	}
	fmt.Println()
}
