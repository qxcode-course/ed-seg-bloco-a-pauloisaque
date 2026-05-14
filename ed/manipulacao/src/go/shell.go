package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	vetor := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			vetor = append(vetor, vet[i])
		}
	}
	return vetor
}

func getCalmWomen(vet []int) []int {
	vetor := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] > -10 && vet[i] < 0 {
			vetor = append(vetor, vet[i])
		}
	}
	return vetor
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if vet[i] < vet[j] {
				k := vet[i]
				l := vet[j]
				vet[i] = l
				vet[j] = k
			}
		}
	}
	return vet
}

func sortStress(vet []int) []int {

	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if math.Abs(float64(vet[i])) < math.Abs(float64(vet[j])) {
				k := vet[i]
				l := vet[j]
				vet[i] = l
				vet[j] = k
			}
		}
	}
	return vet
}

func reverse(vet []int) []int {
	novoVetor := make([]int, 0)
	for i := len(vet) - 1; i >= 0; i-- {
		novoVetor = append(novoVetor, vet[i])
	}
	return novoVetor
}

func unique(vet []int) []int {
	novoVetor := make([]int, 0)

	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if i == j {
				continue
			}
			if vet[i] == 0 {
				continue
			}
			if vet[i] == vet[j] {
				vet[j] = 0
			}
		}
	}

	for i := 0; i < len(vet); i++ {
		if vet[i] != 0 {
			novoVetor = append(novoVetor, vet[i])
		}
	}

	return novoVetor
}

func repeated(vet []int) []int {
	novoVetor := make([]int, 0)

	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if i == j {
				continue
			}
			if vet[i] == 0 {
				continue
			}
			if vet[i] == vet[j] {
				novoVetor = append(novoVetor, vet[j])
				vet[j] = 0
			}
		}
	}

	return sortVet(novoVetor)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
