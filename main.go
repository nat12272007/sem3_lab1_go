package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arr = NewDynamicArray()
var flist = NewForwardList()
var dlist = NewDoublyList()
var stk = NewStack()
var q = NewQueue()

func saveToFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "ARRAY %d\n", arr.GetSize())
	for i := 0; i < arr.GetSize(); i++ {
		fmt.Fprintf(w, "%s ", arr.GetElement(i))
	}

	fmt.Fprintf(w, "\nFLIST %d\n", flist.GetCount())
	for cur := flist.GetHead(); cur != nil; cur = cur.next {
		fmt.Fprintf(w, "%s ", cur.value)
	}

	fmt.Fprintf(w, "\nDLIST %d\n", dlist.GetCount())
	for cur := dlist.GetHead(); cur != nil; cur = cur.next {
		fmt.Fprintf(w, "%s ", cur.value)
	}

	fmt.Fprintf(w, "\nSTACK %d\n", stk.GetTop()+1)
	for i := 0; i <= stk.GetTop(); i++ {
		fmt.Fprintf(w, "%s ", stk.GetElement(i))
	}

	fmt.Fprintf(w, "\nQUEUE %d\n", q.GetCount())
	for i := 0; i < q.GetCount(); i++ {
		fmt.Fprintf(w, "%s ", q.GetElement(i))
	}
	
	fmt.Fprintf(w, "\n")
}

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
			if n < 0 || n > 1000000 { n = 0 }
			arr.Clear()
			for i := 0; i < n; i++ {
				var v string
				fmt.Fscan(r, &v)
				arr.PushBack(v)
			}
		case "FLIST":
			if n < 0 || n > 1000000 { n = 0 }
			flist.Clear()
			for i := 0; i < n; i++ {
				var v string
				fmt.Fscan(r, &v)
				flist.PushBack(v)
			}
		case "DLIST":
			if n < 0 || n > 1000000 { n = 0 }
			dlist.Clear()
			for i := 0; i < n; i++ {
				var v string
				fmt.Fscan(r, &v)
				dlist.PushBack(v)
			}
		case "STACK":
			if n < 0 || n > 1000000 { n = 0 }
			stk.Clear()
			for i := 0; i < n; i++ {
				var v string
				fmt.Fscan(r, &v)
				stk.Push(v)
			}
		case "QUEUE":
			if n < 0 || n > 1000000 { n = 0 }
			q.Clear()
			for i := 0; i < n; i++ {
				var v string
				fmt.Fscan(r, &v)
				q.Push(v)
			}
		}
	}
}

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
			arr.Print()
		case "mylist":
			flist.Print()
		case "mydlist":
			dlist.Print()
		case "mystack":
			stk.Print()
		case "myqueue":
			q.Print()
		}
		return
	}

	switch cmd {
	case "MPUSH":
		arr.PushBack(argStr(fields, 2))
	case "MINSERT":
		arr.PushAt(argInt(fields, 2), argStr(fields, 3))
	case "MGET":
		fmt.Println(arr.Get(argInt(fields, 2)))
	case "MSET":
		arr.Set(argInt(fields, 2), argStr(fields, 3))
	case "MDEL":
		arr.RemoveAt(argInt(fields, 2))
	case "MLEN":
		fmt.Println(arr.Length())

	case "FADDHEAD":
		flist.PushFront(argStr(fields, 2))
	case "FADDTAIL":
		flist.PushBack(argStr(fields, 2))
	case "FADDAFTER":
		t := argStr(fields, 2)
		v := argStr(fields, 3)
		if n := flist.Find(t); n != nil {
			flist.InsertAfter(n, v)
		}
	case "FADDBEFORE":
		t := argStr(fields, 2)
		v := argStr(fields, 3)
		if n := flist.Find(t); n != nil {
			flist.InsertBefore(n, v)
		}
	case "FDELHEAD":
		flist.PopFront()
	case "FDELTAIL":
		flist.PopBack()
	case "FDELAFTER":
		t := argStr(fields, 2)
		if n := flist.Find(t); n != nil {
			flist.RemoveAfter(n)
		}
	case "FDELBEFORE":
		t := argStr(fields, 2)
		if n := flist.Find(t); n != nil {
			flist.RemoveBefore(n)
		}
	case "FDELVAL":
		flist.RemoveByValue(argStr(fields, 2))
	case "FGET":
		fmt.Println(flist.GetByIndex(argInt(fields, 2)))
	case "FFIND":
		if flist.Find(argStr(fields, 2)) != nil {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}

	case "LADDHEAD":
		dlist.PushFront(argStr(fields, 2))
	case "LADDTAIL":
		dlist.PushBack(argStr(fields, 2))
	case "LADDAFTER":
		t := argStr(fields, 2)
		v := argStr(fields, 3)
		if n := dlist.Find(t); n != nil {
			dlist.InsertAfter(n, v)
		}
	case "LADDBEFORE":
		t := argStr(fields, 2)
		v := argStr(fields, 3)
		if n := dlist.Find(t); n != nil {
			dlist.InsertBefore(n, v)
		}
	case "LDELHEAD":
		dlist.PopFront()
	case "LDELTAIL":
		dlist.PopBack()
	case "LDELAFTER":
		t := argStr(fields, 2)
		if n := dlist.Find(t); n != nil && n.next != nil {
			dlist.RemoveNode(n.next)
		}
	case "LDELBEFORE":
		t := argStr(fields, 2)
		if n := dlist.Find(t); n != nil && n.prev != nil {
			dlist.RemoveNode(n.prev)
		}
	case "LDELVAL":
		dlist.RemoveByValue(argStr(fields, 2))
	case "LGET":
		fmt.Println(dlist.GetByIndex(argInt(fields, 2)))
	case "LFIND":
		if dlist.Find(argStr(fields, 2)) != nil {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
	case "LPRINTBACK":
		dlist.PrintBack()

	case "SPUSH":
		stk.Push(argStr(fields, 2))
	case "SPOP":
		v := stk.Peek()
		if stk.Pop() {
			fmt.Println(v)
		}
	case "SPEEK":
		fmt.Println(stk.Peek())

	case "QPUSH":
		q.Push(argStr(fields, 2))
	case "QPOP":
		v := q.Peek()
		if q.Pop() {
			fmt.Println(v)
		}
	case "QPEEK":
		fmt.Println(q.Peek())
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
