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

func NewMultiSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
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

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

func (v *Vector) Insert(value int) {
	if v.data[0] == 0 {
		v.data[0] = value
		v.size++
		return
	}
	for i := 0; i < len(v.data); i++ {
		if v.data[i] == 0 {
			v.data[i] = value
			v.size++
			v.OrganizarPorOrdem()
			return
		}
	}
	v.data = append(v.data, value)
	v.size++
	v.OrganizarPorOrdem()
}

func (v *Vector) OrganizarPorOrdem() {
	for i := 0; i < v.size; i++ {
		for j := 0; j < v.size; j++ {
			if v.data[i] < v.data[j] {
				buf := v.data[j]
				v.data[j] = v.data[i]
				v.data[i] = buf
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

func (v *Vector) Erase(value int) error {
	if v.Contains(value) {
		for i := 0; i < len(v.data); i++ {
			if v.data[i] == value {
				for j := i; j < len(v.data)-1; j++ {
					v.data[j] = v.data[j+1]
				}
				v.data = v.data[:v.size-1]
				v.size--
				return nil
			}
		}
		v.data = v.data[:v.size-1]
		return nil
	}
	return errors.New("value not found")
}

func (v *Vector) Count(value int) int {
	qnt := 0
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			qnt++
		}
	}
	return qnt
}

func (v *Vector) Unique() int {
	unique := 0
	valores := make([]int, 0)
	element := v.data[0]
	valores = append(valores, v.data[0])

	for i := 0; i < v.size; i++ {
		if v.data[i] != element {
			valores = append(valores, v.data[i])
			element = v.data[i]
		}
	}

	return unique
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *Vector = NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
