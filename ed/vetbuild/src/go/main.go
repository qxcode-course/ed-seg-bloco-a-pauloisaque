package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) PushBack(value int) {
	for i := 0; i < len(v.data); i++ {
		if v.data[i] == 0 {
			v.data[i] = value
			v.size++
			if v.size > v.capacity {
				v.capacity = v.capacity * 2
			}
			return
		}
	}
	v.size++
	if v.size > v.capacity {
		v.capacity = v.capacity * 2
	}
	v.data = append(v.data, value)
}

func (v *Vector) PopBack() error {
	if v.size != 0 {
		v.data = v.data[0 : v.size-1]
		v.size--
		return nil
	}
	return errors.New("vector is empty")
}

func (v *Vector) Insert(index int, value int) error {
	v.PushBack(1)
	for i := v.size - 1; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	return nil
}

func (v *Vector) Reserve(newCapacity int) {
	v.capacity = newCapacity
}

func (v *Vector) Erase(index int) error {
	if index < v.size {
		for i := 0; i < v.size; i++ {
			if i == index {
				for j := index + 1; j < v.size; j++ {
					v.data[j-1] = v.data[j]
				}
			}
		}
		v.PopBack()
		return nil
	}
	return errors.New("index out of range")
}

func (v *Vector) Clear() {
	v.size = 0
	v.data = v.data[:]
}

func (v *Vector) At(value int) (int, error) {
	if value <= v.size {
		return v.data[value], nil
	}
	return 0, errors.New("index out of range")
}

func (v *Vector) Set(index int, value int) error {
	if index < v.size {
		v.data[index] = value
		return nil
	}
	return errors.New("index out of range")
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

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

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var v *Vector = NewVector(0)
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
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println("index out of range")
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println("index out of range")
			}
			//
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
