package main

import (
	"errors"
	"fmt"
)

type Character rune

type List struct {
	elements []Character
}

func NewList() *List {
	return &List{elements: make([]Character, 0)}
}

func (list *List) Length() int {
	return len(list.elements)
}

func (list *List) Append(element Character) {
	list.elements = append(list.elements, element)
}

func (list *List) Insert(element Character, index int) error {
	if index < 0 || index > len(list.elements) {
		return errors.New("індекс поза межами діапазону")
	}

	list.elements = append(list.elements, 0)
	copy(list.elements[index+1:], list.elements[index:])
	list.elements[index] = element
	return nil
}

func (list *List) Delete(index int) (Character, error) {
	if index < 0 || index >= len(list.elements) {
		return 0, errors.New("індекс поза межами діапазону")
	}

	value := list.elements[index]
	list.elements = append(list.elements[:index], list.elements[index+1:]...)
	return value, nil
}

func (list *List) DeleteAll(element Character) {
	newElements := make([]Character, 0)
	for _, v := range list.elements {
		if v != element {
			newElements = append(newElements, v)
		}
	}
	list.elements = newElements
}

func (list *List) Get(index int) (Character, error) {
	if index < 0 || index >= len(list.elements) {
		return 0, errors.New("індекс поза межами діапазону")
	}
	return list.elements[index], nil
}

func (list *List) Clone() *List {
	newList := NewList()
	newList.elements = append(newList.elements, list.elements...)
	return newList
}

func (list *List) Reverse() {
	for i, j := 0, len(list.elements)-1; i < j; i, j = i+1, j-1 {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List) FindFirst(element Character) int {
	for i, v := range list.elements {
		if v == element {
			return i
		}
	}
	return -1
}

func (list *List) FindLast(element Character) int {
	for i := len(list.elements) - 1; i >= 0; i-- {
		if list.elements[i] == element {
			return i
		}
	}
	return -1
}

func (list *List) Clear() {
	list.elements = make([]Character, 0)
}

func (list *List) Extend(other *List) {
	list.elements = append(list.elements, other.elements...)
}

func (list *List) Print() {
	for _, v := range list.elements {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
