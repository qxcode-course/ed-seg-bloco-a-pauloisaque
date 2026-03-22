package main

import "fmt"

type Jogada struct {
	pa, pb int
}

func main() {
	qtd := 0
	fmt.Scan(&qtd)
	jogadas := make([]Jogada, qtd)
	for i := range jogadas {
		fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
	}
	fmt.Println(jogar(jogadas))
}

func jogar(jogadas []Jogada) int32 {
	var placarFinal [][]int32
	fmt.Println(placarFinal)
	for i := range jogadas {
		var resultadoJogada []int32
		resultadoJogada = append(resultadoJogada, int32(jogadas[i].pb-jogadas[i].pa), int32(i))
		placarFinal = append(placarFinal, resultadoJogada)
	}
	return 0
}
