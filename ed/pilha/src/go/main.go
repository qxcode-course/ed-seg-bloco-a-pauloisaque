package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data []T
	size int
}

func NewStack[T any](value int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, value),
		size: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	if len(s.data) == s.size {
		aux := make([]T, s.size*2)
		for i := 0; i < s.size; i++ {
			aux[i] = s.data[i]
		}
		s.data = aux
	}
	s.data[s.size] = value
	s.size++
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Peek() (T, error) {
	if s.size == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	resul := s.data[s.size-1]
	return resul, nil
}

func (s *Stack[T]) Clear() {
	s.size = 0
}

func (s *Stack[T]) Pop() error {
	if s.size == 0 {
		return errors.New("stack is empty")
	}
	aux := make([]T, len(s.data))
	for i := 0; i < s.size; i++ {
		aux[i] = s.data[i]
	}
	s.data = aux
	s.size--
	return nil
}

func (s *Stack[T]) String() string {
	output := ""
	for i := range cap(s.data) {
		if i != 0 {
			output += ", "
		}
		if i < s.size {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			v = NewStack[int](cap)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v)
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
