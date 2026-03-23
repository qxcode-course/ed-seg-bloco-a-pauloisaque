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
	var resultadoJogadas = jogar(jogadas)
	if resultadoJogadas == 0 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(jogar(jogadas))
	}
}

func jogar(jogadas []Jogada) int32 {
	var placarFinal [][]int32
	var vencedor int32

	for i := range jogadas {
		if jogadas[i].pa < 10 || jogadas[i].pb < 10 {
			placarFinal = append(placarFinal, []int32{100, int32(i)})
			continue
		}
		var resultadoJogada []int32
		resultadoJogada = append(resultadoJogada, int32(jogadas[i].pb-jogadas[i].pa), int32(i))
		placarFinal = append(placarFinal, resultadoJogada)
	}

	var minimo int32 = placarFinal[0][0]
	for _, v := range placarFinal {
		if v[0] < minimo {
			minimo = v[0]
			vencedor = v[1]
		}
	}
	return vencedor
}
