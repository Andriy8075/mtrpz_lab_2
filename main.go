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

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Length() int {
	return dll.size
}

func (dll *DoublyLinkedList) Append(element Character) {
	newNode := &Node{value: element, prev: dll.tail}

	if dll.tail == nil {
		dll.head = newNode
	} else {
		dll.tail.next = newNode
	}

	dll.tail = newNode
	dll.size++
}

func (dll *DoublyLinkedList) Insert(element Character, index int) error {
	if index < 0 || index > dll.size {
		return errors.New("індекс поза межами діапазону")
	}

	if index == dll.size {
		dll.Append(element)
		return nil
	}

	newNode := &Node{value: element}
	if index == 0 {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	} else {
		current := dll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}

	dll.size++
	return nil
}

func (dll *DoublyLinkedList) Delete(index int) (Character, error) {
	if index < 0 || index >= dll.size {
		return 0, errors.New("індекс поза межами діапазону")
	}

	var value Character
	if dll.size == 1 {
		value = dll.head.value
		dll.head = nil
		dll.tail = nil
	} else if index == 0 {
		value = dll.head.value
		dll.head = dll.head.next
		dll.head.prev = nil
	} else if index == dll.size-1 {
		value = dll.tail.value
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	} else {
		current := dll.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		value = current.value
		current.prev.next = current.next
		current.next.prev = current.prev
	}

	dll.size--
	return value, nil
}

func (dll *DoublyLinkedList) DeleteAll(element Character) {
	current := dll.head
	for current != nil {
		if current.value == element {
			if current.prev == nil {
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				}
			} else {
				current.prev.next = current.next
			}

			if current.next == nil {
				dll.tail = current.prev
				if dll.tail != nil {
					dll.tail.next = nil
				}
			} else {
				current.next.prev = current.prev
			}

			dll.size--
		}
		current = current.next
	}
}

func (dll *DoublyLinkedList) Get(index int) (Character, error) {
	if index < 0 || index >= dll.size {
		return 0, errors.New("індекс поза межами діапазону")
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (dll *DoublyLinkedList) Clone() *DoublyLinkedList {
	newList := NewDoublyLinkedList()
	current := dll.head
	for current != nil {
		newList.Append(current.value)
		current = current.next
	}
	return newList
}

func (dll *DoublyLinkedList) Reverse() {
	current := dll.head
	dll.tail = current
	var temp *Node

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		dll.head = temp.prev

	}
}

func (dll *DoublyLinkedList) FindFirst(element Character) int {
	current := dll.head
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

func (dll *DoublyLinkedList) FindLast(element Character) int {
	current := dll.tail
	index := dll.size - 1
	for current != nil {
		if current.value == element {
			return index
		}
		current = current.prev
		index--
	}

	return -1
}

func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

func (dll *DoublyLinkedList) Extend(other *DoublyLinkedList) {
	current := other.head
	for current != nil {
		dll.Append(current.value)
		current = current.next
	}
}

func (dll *DoublyLinkedList) Print() {
	current := dll.head
	for current != nil {
		fmt.Printf("%c ", current.value)
		current = current.next
	}
	fmt.Println()
}
