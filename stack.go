package main

import "fmt"

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{data: make([]string, 0, 4)}
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() bool {
	if len(s.data) == 0 {
		return false
	}
	s.data = s.data[:len(s.data)-1]
	return true
}

func (s *Stack) Peek() string {
	if len(s.data) == 0 {
		return "ERROR"
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Print() {
	fmt.Print("Stack(top->bottom): [")
	for i := len(s.data) - 1; i >= 0; i-- {
		if i != len(s.data)-1 {
			fmt.Print(", ")
		}
		fmt.Print(s.data[i])
	}
	fmt.Println("]")
}

func (s *Stack) Clear() {
	s.data = make([]string, 0, 4)
}

func (s *Stack) GetTop() int {
	return len(s.data) - 1
}

func (s *Stack) GetElement(index int) string {
	return s.data[index]
}
