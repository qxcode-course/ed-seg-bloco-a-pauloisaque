package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.data)
}

type List struct {
	head *Node
}

func (l *List) String() string {
	result := "["
	for x := l.head.next; x != l.head; x = x.next {
		if x == l.head.next {
			result = result + x.String()
			continue
		}
		result = result + ", " + x.String()
	}
	result = result + "]"
	return result
}

func NewList() *List {
	list := List{}
	list.head = &Node{}
	list.head.next = list.head
	list.head.prev = list.head
	return &list
}

func inserir(node *Node, value int) {
	novo := &Node{
		data: value,
		next: node,
		prev: node.prev,
	}
	node.prev.next = novo
	node.prev = novo
}

func (l *List) PushFront(value int) {
	inserir(l.head.next, value)
}

func (l *List) PushBack(value int) {
	inserir(l.head, value)
}

func (l *List) Size() int {
	size := 0
	for x := l.head.next; x != l.head; x = x.next {
		size++
	}
	return size
}

func (l *List) Clear() {
	l.head.next = l.head
	l.head.prev = l.head
}

func (l *List) PopFront() {
	l.head.next = l.head.next.next
	l.head.next.prev = l.head
}

func (l *List) PopBack() {
	l.head.prev = l.head.prev.prev
	l.head.prev.next = l.head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
