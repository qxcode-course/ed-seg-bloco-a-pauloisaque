package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Insert(value int) {
	if v.size == 0 {
		v.size++
		v.data[0] = value
		return
	}
	for i := 0; i < len(v.data); i++ {
		if v.data[i] == value {
			return
		}
		if v.data[i] == 0 {
			v.data[i] = value
			v.size++
			v.organizaVetor()
			return
		}
	}
	v.data = append(v.data, value)
	v.size++
	v.organizaVetor()
}

func (v *Vector) organizaVetor() {
	for i := 0; i < v.size; i++ {
		for j := 0; j < v.size; j++ {
			if v.data[i] < v.data[j] {
				buffer := v.data[j]
				v.data[j] = v.data[i]
				v.data[i] = buffer
			}
		}
	}
}

func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return true
		}
	}
	return false
}

func (v *Vector) Erase(value int) bool {
	if v.Contains(value) {
		for i := 0; i < v.size; i++ {
			if v.data[i] == value {
				for j := i; j < v.size-1; j++ {
					v.data[j] = v.data[j+1]
				}
			}
		}
		v.data = v.data[:v.size-1]
		v.size--
		return true
	}
	fmt.Println("value not found")
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var v *Vector = NewSet(0)

	for scanner.Scan() {
		fmt.Print("$")
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
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
