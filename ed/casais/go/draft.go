package main

import "fmt"

func main() {
	var qnt int
	fmt.Scan(&qnt)
	vetor := make([]int, qnt)
	for i := 0; i < qnt; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Print(verificaPares(vetor, qnt))

}

func verificaPares(vetor []int, qnt int) int {
	novoVetor := append([]int{}, vetor...)
	pontaUm := 0
	pontaDois := 0
	countPares := 0

	iniciarProcesso := func() {
		pontaUm = 0
		pontaDois = 0
	}

	for pontaDois < qnt {
		if novoVetor[pontaUm] == novoVetor[pontaDois] && pontaUm != pontaDois {
			novoVetor = append(vetor[:pontaUm], vetor[pontaUm+1:]...)
			novoVetor = append(vetor[:pontaDois], vetor[pontaDois+1:]...)
			countPares++
			iniciarProcesso()
		} else {
			if pontaDois > qnt {
				break
			}
			pontaDois++
		}
	}

	return countPares

}
