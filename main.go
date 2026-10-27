package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//a.массив
type DynamicArray struct {
	data []int
}

func (a *DynamicArray) init(cap int) {
	a.data = make([]int, 0, cap)
}

func (a *DynamicArray) pushBack(value int) {
	a.data = append(a.data, value)
}

func (a *DynamicArray) pushAt(index, value int) {
	if index < 0 || index > len(a.data) {
		return
	}
	a.data = append(a.data, 0)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value
}

func (a *DynamicArray) get(index int) int {
	if index >= 0 && index < len(a.data) {
		return a.data[index]
	}
	return -1
}

func (a *DynamicArray) removeAt(index int) {
	if index < 0 || index >= len(a.data) {
		return
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
}

func (a *DynamicArray) set(index, value int) {
	if index >= 0 && index < len(a.data) {
		a.data[index] = value
	}
}

func (a *DynamicArray) length() int {
	return len(a.data)
}

func (a *DynamicArray) print() {
	fmt.Print("[")
	for i, v := range a.data {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func (a *DynamicArray) clear() {
	a.data = nil
}

//b.односвязный список
type FNode struct {
	value int
	next  *FNode
}

type ForwardList struct {
	head  *FNode
	tail  *FNode
	count int
}

func (l *ForwardList) pushFront(value int) {
	node := &FNode{value: value, next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.count++
}

func (l *ForwardList) pushBack(value int) {
	node := &FNode{value: value, next: nil}
	if l.tail != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.count++
}

func (l *ForwardList) insertAfter(prev *FNode, value int) {
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

func (l *ForwardList) insertBefore(target *FNode, value int) {
	if target == nil {
		return
	}
	if l.head == target {
		l.pushFront(value)
		return
	}
	prev := l.head
	for prev != nil && prev.next != target {
		prev = prev.next
	}
	if prev != nil {
		l.insertAfter(prev, value)
	}
}

func (l *ForwardList) popFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.count--
}

func (l *ForwardList) popBack() {
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

func (l *ForwardList) removeAfter(prev *FNode) {
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

func (l *ForwardList) removeBefore(target *FNode) {
	if l.head == nil || l.head == target {
		return
	}
	if l.head.next == target {
		l.popFront()
		return
	}
	prev := l.head
	for prev.next != nil && prev.next.next != target {
		prev = prev.next
	}
	if prev.next != nil {
		l.removeAfter(prev)
	}
}

func (l *ForwardList) find(value int) *FNode {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == value {
			return cur
		}
	}
	return nil
}

func (l *ForwardList) removeByValue(value int) bool {
	if l.head == nil {
		return false
	}
	if l.head.value == value {
		l.popFront()
		return true
	}
	prev := l.head
	for prev.next != nil && prev.next.value != value {
		prev = prev.next
	}
	if prev.next != nil {
		l.removeAfter(prev)
		return true
	}
	return false
}

func (l *ForwardList) getByIndex(index int) int {
	if index < 0 {
		return -1
	}
	cur := l.head
	for i := 0; cur != nil && i < index; i++ {
		cur = cur.next
	}
	if cur != nil {
		return cur.value
	}
	return -1
}

func (l *ForwardList) print() {
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

func (l *ForwardList) clear() {
	for l.head != nil {
		l.popFront()
	}
}

//c.двусвязный список
type LNode struct {
	value int
	prev  *LNode
	next  *LNode
}

type DoublyList struct {
	head  *LNode
	tail  *LNode
	count int
}

func (l *DoublyList) pushFront(value int) {
	node := &LNode{value: value, prev: nil, next: l.head}
	if l.head != nil {
		l.head.prev = node
	} else {
		l.tail = node
	}
	l.head = node
	l.count++
}

func (l *DoublyList) pushBack(value int) {
	node := &LNode{value: value, prev: l.tail, next: nil}
	if l.tail != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.count++
}

func (l *DoublyList) insertAfter(pos *LNode, value int) {
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

func (l *DoublyList) insertBefore(pos *LNode, value int) {
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

func (l *DoublyList) popFront() {
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

func (l *DoublyList) popBack() {
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

func (l *DoublyList) removeNode(node *LNode) {
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

func (l *DoublyList) find(value int) *LNode {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == value {
			return cur
		}
	}
	return nil
}

func (l *DoublyList) removeByValue(value int) bool {
	node := l.find(value)
	if node != nil {
		l.removeNode(node)
		return true
	}
	return false
}

func (l *DoublyList) getByIndex(index int) int {
	if index < 0 {
		return -1
	}
	cur := l.head
	for i := 0; cur != nil && i < index; i++ {
		cur = cur.next
	}
	if cur != nil {
		return cur.value
	}
	return -1
}

func (l *DoublyList) print() {
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

func (l *DoublyList) printBack() {
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

func (l *DoublyList) clear() {
	for l.head != nil {
		l.popFront()
	}
}

//d.стек
type Stack struct {
	data []int
}

func (s *Stack) init(cap int) {
	s.data = make([]int, 0, cap)
}

func (s *Stack) push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) pop() bool {
	if len(s.data) == 0 {
		return false
	}
	s.data = s.data[:len(s.data)-1]
	return true
}

func (s *Stack) peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) print() {
	fmt.Print("Stack(top->bottom): [")
	for i := len(s.data) - 1; i >= 0; i-- {
		if i != len(s.data)-1 {
			fmt.Print(", ")
		}
		fmt.Print(s.data[i])
	}
	fmt.Println("]")
}

func (s *Stack) clear() {
	s.data = nil
}

//e.очередь
type Queue struct {
	data []int
}

func (q *Queue) init(cap int) {
	q.data = make([]int, 0, cap)
}

func (q *Queue) push(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) pop() bool {
	if len(q.data) == 0 {
		return false
	}
	q.data = q.data[1:]
	return true
}

func (q *Queue) peek() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) print() {
	fmt.Print("Queue(front->back): [")
	for i, v := range q.data {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func (q *Queue) clear() {
	q.data = nil
}

//глобальные экземпляры
var arr DynamicArray
var flist ForwardList
var dlist DoublyList
var stk Stack
var q Queue

//утилиты
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//сохранение
func saveToFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "ARRAY %d\n", len(arr.data))
	for _, v := range arr.data {
		fmt.Fprintf(w, "%d ", v)
	}

	fmt.Fprintf(w, "\nFLIST %d\n", flist.count)
	for cur := flist.head; cur != nil; cur = cur.next {
		fmt.Fprintf(w, "%d ", cur.value)
	}

	fmt.Fprintf(w, "\nDLIST %d\n", dlist.count)
	for cur := dlist.head; cur != nil; cur = cur.next {
		fmt.Fprintf(w, "%d ", cur.value)
	}

	fmt.Fprintf(w, "\nSTACK %d\n", len(stk.data))
	for _, v := range stk.data {
		fmt.Fprintf(w, "%d ", v)
	}

	fmt.Fprintf(w, "\nQUEUE %d\n", len(q.data))
	for _, v := range q.data {
		fmt.Fprintf(w, "%d ", v)
	}

	fmt.Fprintf(w, "\n")
}

//загрузка
func loadFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("[File not found, starting empty: %s]\n", path)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		var typ string
		var n int

		if _, err := fmt.Fscan(r, &typ, &n); err != nil {
			break
		}

		switch typ {
		case "ARRAY":
			if n < 0 || n > 1000000 {
				n = 0
			}
			arr.init(maxInt(4, n))
			arr.data = make([]int, n)
			for i := 0; i < n; i++ {
				fmt.Fscan(r, &arr.data[i])
			}

		case "FLIST":
			if n < 0 || n > 1000000 {
				n = 0
			}
			flist.clear()
			for i := 0; i < n; i++ {
				var v int
				fmt.Fscan(r, &v)
				flist.pushBack(v)
			}

		case "DLIST":
			if n < 0 || n > 1000000 {
				n = 0
			}
			dlist.clear()
			for i := 0; i < n; i++ {
				var v int
				fmt.Fscan(r, &v)
				dlist.pushBack(v)
			}

		case "STACK":
			if n < 0 || n > 1000000 {
				n = 0
			}
			stk.clear()
			stk.init(maxInt(4, n))
			for i := 0; i < n; i++ {
				var v int
				fmt.Fscan(r, &v)
				stk.push(v)
			}

		case "QUEUE":
			if n < 0 || n > 1000000 {
				n = 0
			}
			q.clear()
			q.init(maxInt(4, n))
			for i := 0; i < n; i++ {
				var v int
				fmt.Fscan(r, &v)
				q.push(v)
			}
		}
	}
}

//парсинг аргументов
func argStr(fields []string, i int) string {
	if i >= 0 && i < len(fields) {
		return fields[i]
	}
	return ""
}

func argInt(fields []string, i int) int {
	if i >= 0 && i < len(fields) {
		v, _ := strconv.Atoi(fields[i])
		return v
	}
	return 0
}

//обработка команд
func processCommand(query string) {
	fields := strings.Fields(query)
	if len(fields) == 0 {
		return
	}

	cmd := fields[0]

	if cmd == "PRINT" {
		target := argStr(fields, 1)
		switch target {
		case "myarr":
			arr.print()
		case "mylist":
			flist.print()
		case "mydlist":
			dlist.print()
		case "mystack":
			stk.print()
		case "myqueue":
			q.print()
		}
		return
	}

	switch cmd {
	//массив
	case "MPUSH":
		arr.pushBack(argInt(fields, 2))
	case "MINSERT":
		arr.pushAt(argInt(fields, 2), argInt(fields, 3))
	case "MGET":
		fmt.Println(arr.get(argInt(fields, 2)))
	case "MSET":
		arr.set(argInt(fields, 2), argInt(fields, 3))
	case "MDEL":
		arr.removeAt(argInt(fields, 2))
	case "MLEN":
		fmt.Println(arr.length())

	//односвязный список
	case "FADDHEAD":
		flist.pushFront(argInt(fields, 2))
	case "FADDTAIL":
		flist.pushBack(argInt(fields, 2))
	case "FADDAFTER":
		t := argInt(fields, 2)
		v := argInt(fields, 3)
		if n := flist.find(t); n != nil {
			flist.insertAfter(n, v)
		}
	case "FADDBEFORE":
		t := argInt(fields, 2)
		v := argInt(fields, 3)
		if n := flist.find(t); n != nil {
			flist.insertBefore(n, v)
		}
	case "FDELHEAD":
		flist.popFront()
	case "FDELTAIL":
		flist.popBack()
	case "FDELAFTER":
		t := argInt(fields, 2)
		if n := flist.find(t); n != nil {
			flist.removeAfter(n)
		}
	case "FDELBEFORE":
		t := argInt(fields, 2)
		if n := flist.find(t); n != nil {
			flist.removeBefore(n)
		}
	case "FDELVAL":
		flist.removeByValue(argInt(fields, 2))
	case "FGET":
		fmt.Println(flist.getByIndex(argInt(fields, 2)))
	case "FFIND":
		if flist.find(argInt(fields, 2)) != nil {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}

	//двусвязный список
	case "LADDHEAD":
		dlist.pushFront(argInt(fields, 2))
	case "LADDTAIL":
		dlist.pushBack(argInt(fields, 2))
	case "LADDAFTER":
		t := argInt(fields, 2)
		v := argInt(fields, 3)
		if n := dlist.find(t); n != nil {
			dlist.insertAfter(n, v)
		}
	case "LADDBEFORE":
		t := argInt(fields, 2)
		v := argInt(fields, 3)
		if n := dlist.find(t); n != nil {
			dlist.insertBefore(n, v)
		}
	case "LDELHEAD":
		dlist.popFront()
	case "LDELTAIL":
		dlist.popBack()
	case "LDELAFTER":
		t := argInt(fields, 2)
		if n := dlist.find(t); n != nil && n.next != nil {
			dlist.removeNode(n.next)
		}
	case "LDELBEFORE":
		t := argInt(fields, 2)
		if n := dlist.find(t); n != nil && n.prev != nil {
			dlist.removeNode(n.prev)
		}
	case "LDELVAL":
		dlist.removeByValue(argInt(fields, 2))
	case "LGET":
		fmt.Println(dlist.getByIndex(argInt(fields, 2)))
	case "LFIND":
		if dlist.find(argInt(fields, 2)) != nil {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
	case "LPRINTBACK":
		dlist.printBack()

	//стек
	case "SPUSH":
		stk.push(argInt(fields, 2))
	case "SPOP":
		v := stk.peek()
		if stk.pop() {
			fmt.Println(v)
		}
	case "SPEEK":
		fmt.Println(stk.peek())

	//очередь
	case "QPUSH":
		q.push(argInt(fields, 2))
	case "QPOP":
		v := q.peek()
		if q.pop() {
			fmt.Println(v)
		}
	case "QPEEK":
		fmt.Println(q.peek())
	}
}

func main() {
	filePath := ""
	var queries []string
	replMode := false

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file":
			if i+1 < len(args) {
				filePath = args[i+1]
				i++
			}
		case "--query":
			if i+1 < len(args) {
				queries = append(queries, args[i+1])
				i++
			}
		case "--repl":
			replMode = true
		}
	}

	if filePath == "" {
		filePath = "data.txt"
		fmt.Println("[No --file given, using data.txt]")
	}

	arr.init(4)
	stk.init(4)
	q.init(4)

	loadFromFile(filePath)

	if len(queries) > 0 {
		for _, qr := range queries {
			processCommand(qr)
		}
		saveToFile(filePath)
	} else if replMode {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "EXIT" || line == "exit" {
				break
			}
			if line != "" {
				processCommand(line)
			}
		}
		saveToFile(filePath)
	}

	fmt.Print("\n[Программа завершена. Нажмите Enter для выхода.]")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
