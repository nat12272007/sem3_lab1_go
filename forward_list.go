package main

import "fmt"

type FNode struct {
	value string
	next  *FNode
}

type ForwardList struct {
	head  *FNode
	tail  *FNode
	count int
}

func NewForwardList() *ForwardList {
	return &ForwardList{}
}

func (l *ForwardList) PushFront(value string) {
	node := &FNode{value: value, next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.count++
}

func (l *ForwardList) PushBack(value string) {
	node := &FNode{value: value, next: nil}
	if l.tail != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.count++
}

func (l *ForwardList) InsertAfter(prev *FNode, value string) {
	if prev == nil {
		return
	}
	node := &FNode{value: value, next: prev.next}
	prev.next = node
	if prev == l.tail {
		l.tail = node
	}
	l.count++
}

func (l *ForwardList) InsertBefore(target *FNode, value string) {
	if target == nil {
		return
	}
	if l.head == target {
		l.PushFront(value)
		return
	}
	prev := l.head
	for prev != nil && prev.next != target {
		prev = prev.next
	}
	if prev != nil {
		l.InsertAfter(prev, value)
	}
}

func (l *ForwardList) PopFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.count--
}

func (l *ForwardList) PopBack() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.count--
		return
	}
	prev := l.head
	for prev.next != l.tail {
		prev = prev.next
	}
	prev.next = nil
	l.tail = prev
	l.count--
}

func (l *ForwardList) RemoveAfter(prev *FNode) {
	if prev == nil || prev.next == nil {
		return
	}
	toDelete := prev.next
	prev.next = toDelete.next
	if toDelete == l.tail {
		l.tail = prev
	}
	l.count--
}

func (l *ForwardList) RemoveBefore(target *FNode) {
	if l.head == nil || l.head == target {
		return
	}
	if l.head.next == target {
		l.PopFront()
		return
	}
	prev := l.head
	for prev.next != nil && prev.next.next != target {
		prev = prev.next
	}
	if prev.next != nil {
		l.RemoveAfter(prev)
	}
}

func (l *ForwardList) Find(value string) *FNode {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == value {
			return cur
		}
	}
	return nil
}

func (l *ForwardList) RemoveByValue(value string) bool {
	if l.head == nil {
		return false
	}
	if l.head.value == value {
		l.PopFront()
		return true
	}
	prev := l.head
	for prev.next != nil && prev.next.value != value {
		prev = prev.next
	}
	if prev.next != nil {
		l.RemoveAfter(prev)
		return true
	}
	return false
}

func (l *ForwardList) GetByIndex(index int) string {
	if index < 0 {
		return "ERROR"
	}
	cur := l.head
	for i := 0; cur != nil && i < index; i++ {
		cur = cur.next
	}
	if cur != nil {
		return cur.value
	}
	return "ERROR"
}

func (l *ForwardList) Print() {
	fmt.Print("[")
	cur := l.head
	first := true
	for cur != nil {
		if !first {
			fmt.Print(" -> ")
		}
		fmt.Print(cur.value)
		cur = cur.next
		first = false
	}
	fmt.Println("]")
}

func (l *ForwardList) Clear() {
	for l.head != nil {
		l.PopFront()
	}
}

func (l *ForwardList) GetCount() int {
	return l.count
}

func (l *ForwardList) GetHead() *FNode {
	return l.head
}
