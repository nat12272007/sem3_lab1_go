package main

import "fmt"

type DynamicArray struct {
	data []string
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{data: make([]string, 0, 4)}
}

func (a *DynamicArray) PushBack(value string) {
	a.data = append(a.data, value)
}

func (a *DynamicArray) PushAt(index int, value string) {
	if index < 0 || index > len(a.data) {
		return
	}
	a.data = append(a.data, "")
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value
}

func (a *DynamicArray) Get(index int) string {
	if index >= 0 && index < len(a.data) {
		return a.data[index]
	}
	return "ERROR"
}

func (a *DynamicArray) RemoveAt(index int) {
	if index < 0 || index >= len(a.data) {
		return
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
}

func (a *DynamicArray) Set(index int, value string) {
	if index >= 0 && index < len(a.data) {
		a.data[index] = value
	}
}

func (a *DynamicArray) Length() int {
	return len(a.data)
}

func (a *DynamicArray) Print() {
	fmt.Print("[")
	for i, v := range a.data {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func (a *DynamicArray) Clear() {
	a.data = make([]string, 0, 4)
}

func (a *DynamicArray) GetSize() int {
	return len(a.data)
}

func (a *DynamicArray) GetElement(index int) string {
	return a.data[index]
}
