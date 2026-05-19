package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func auxilia_tostr(vet []int) string {
	aux := vet
	if len(vet) == 0 {
		return ""
	}
	if len(vet) == 1 {
		return fmt.Sprint(vet[0])
	}
	vet = vet[1:]
	return fmt.Sprintf("%d, ", aux[0]) + auxilia_tostr(vet)
}

func tostr(vet []int) string {
	return "[" + auxilia_tostr(vet) + "]"
}

func auxilia_tostrrev(vet []int, i int) []int {
	vet = append(vet, vet[0])
	vet = vet[1:]
	if i == 0 {
		return vet
	}
	i--
	return auxilia_tostrrev(vet, i)
}

func tostrrev(vet []int) string {
	return tostr(auxilia_tostrrev(vet, len(vet)-1))
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	_ = vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	return 0
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet
	return 0
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet
	return 0
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
