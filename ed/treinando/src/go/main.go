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

func auxilia_tostrrev(vet []int) []int {
	if len(vet) >= 1 {
		last := vet[len(vet)-1]
		return append([]int{last}, auxilia_tostrrev(vet[:len(vet)-1])...)
	}
	return vet
}

func tostrrev(vet []int) string {
	return tostr(auxilia_tostrrev(vet))
}

// reverse: inverte os elementos do slice
func reverse(vet []int) []int {
	return auxilia_tostrrev(vet)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {

}

func aux_min(vet[int]) int

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
			vet = reverse(vet)
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
