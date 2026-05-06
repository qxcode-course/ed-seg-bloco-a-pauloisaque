package main

import (
	"fmt"
	"math"
)

func main() {
	var qnt int
	fmt.Scan(&qnt)
	vetor := make([]int, qnt)
	for i := 0; i < qnt; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Println(verificaPares(vetor, qnt))

}

func verificaPares(vetor []int, qnt int) int {
	novoVetor := append([]int{}, vetor...)
	pontaUm := 0
	pontaDois := 0
	countPares := 0

	iniciarProcesso := func() {
		pontaDois = 0
	}
	abs := func(x int) int {
		return int(math.Abs(float64(x)))
	}
	for pontaUm < qnt-1 {
		if abs(novoVetor[pontaUm]) == abs(novoVetor[pontaDois]) && novoVetor[pontaUm] != 0 && novoVetor[pontaDois] != 0 && pontaUm != pontaDois {
			novoVetor[pontaDois] = 0
			novoVetor[pontaUm] = 0
			countPares++
			iniciarProcesso()
		} else {
			pontaDois++
			if pontaDois == qnt {
				pontaDois--
				pontaUm++
			}
		}
	}

	return countPares

}
