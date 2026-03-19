package main

import "fmt"

type Jogada struct {
	pa, pb int
}

func main() {
	qtd := 0
	fmt.Scan(&qtd)
	jogadas := make([]Jogada, qtd)
	for _, jog := range jogadas {
		fmt.Scan(&jog.pa, &jog.pb)
	}
	fmt.Println(jogadas)
}
