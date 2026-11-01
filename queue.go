package main

import "fmt"

type Queue struct {
	data []string
}

func NewQueue() *Queue {
	return &Queue{data: make([]string, 0, 4)}
}

func (q *Queue) Push(value string) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() bool {
	if len(q.data) == 0 {
		return false
	}
	q.data = q.data[1:]
	return true
}

func (q *Queue) Peek() string {
	if len(q.data) == 0 {
		return "ERROR"
	}
	return q.data[0]
}

func (q *Queue) Print() {
	fmt.Print("Queue(front->back): [")
	for i, v := range q.data {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func (q *Queue) Clear() {
	q.data = make([]string, 0, 4)
}

func (q *Queue) GetCount() int {
	return len(q.data)
}

func (q *Queue) GetElement(index int) string {
	return q.data[index]
}
