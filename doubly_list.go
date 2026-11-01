package main

import "fmt"

type LNode struct {
	value string
	prev  *LNode
	next  *LNode
}

type DoublyList struct {
	head  *LNode
	tail  *LNode
	count int
}

func NewDoublyList() *DoublyList {
	return &DoublyList{}
}

func (l *DoublyList) PushFront(value string) {
	node := &LNode{value: value, prev: nil, next: l.head}
	if l.head != nil {
		l.head.prev = node
	} else {
		l.tail = node
	}
	l.head = node
	l.count++
}

func (l *DoublyList) PushBack(value string) {
	node := &LNode{value: value, prev: l.tail, next: nil}
	if l.tail != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.count++
}

func (l *DoublyList) InsertAfter(pos *LNode, value string) {
	if pos == nil {
		return
	}
	node := &LNode{value: value, prev: pos, next: pos.next}
	if pos.next != nil {
		pos.next.prev = node
	} else {
		l.tail = node
	}
	pos.next = node
	l.count++
}

func (l *DoublyList) InsertBefore(pos *LNode, value string) {
	if pos == nil {
		return
	}
	node := &LNode{value: value, prev: pos.prev, next: pos}
	if pos.prev != nil {
		pos.prev.next = node
	} else {
		l.head = node
	}
	pos.prev = node
	l.count++
}

func (l *DoublyList) PopFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.count--
}

func (l *DoublyList) PopBack() {
	if l.tail == nil {
		return
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.count--
}

func (l *DoublyList) RemoveNode(node *LNode) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	l.count--
}

func (l *DoublyList) Find(value string) *LNode {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == value {
			return cur
		}
	}
	return nil
}

func (l *DoublyList) RemoveByValue(value string) bool {
	node := l.Find(value)
	if node != nil {
		l.RemoveNode(node)
		return true
	}
	return false
}

func (l *DoublyList) GetByIndex(index int) string {
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

func (l *DoublyList) Print() {
	fmt.Print("[")
	cur := l.head
	first := true
	for cur != nil {
		if !first {
			fmt.Print(" <-> ")
		}
		fmt.Print(cur.value)
		cur = cur.next
		first = false
	}
	fmt.Println("]")
}

func (l *DoublyList) PrintBack() {
	fmt.Print("[")
	cur := l.tail
	first := true
	for cur != nil {
		if !first {
			fmt.Print(" <-> ")
		}
		fmt.Print(cur.value)
		cur = cur.prev
		first = false
	}
	fmt.Println("]")
}

func (l *DoublyList) Clear() {
	for l.head != nil {
		l.PopFront()
	}
}

func (l *DoublyList) GetCount() int {
	return l.count
}

func (l *DoublyList) GetHead() *LNode {
	return l.head
}
