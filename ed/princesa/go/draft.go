package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	//cria a fila inicial com as pessoas de 1 até n
	vivos := make([]int, n)
	for i := 0; i < n; i++ {
		vivos[i] = i + 1
	}

	//a espada começa com a pessoa E
	pos := e - 1
	// 1 2> 3 4
	// 0 1  2 3
	//enquanto tiver mais de uma pessoa viva, a simulação irá continuar
	for len(vivos) > 1 {
		// mostra o estado atual da fila e quem está com a espada
		imprime(vivos, pos)
		// a vitima e a proxima pessoa viva
		vitima := (pos + 1) % len(vivos)
		//remove a vitima do vetor
		vivos = append(vivos[:vitima], vivos[vitima+1:]...)
		//depois da morte, a espada passa a proxima pessoa
		//se a vitima estava no fim do veto, a espada volta ao inicio
		pos = vitima % len(vivos)
	}
	//imprime o vencedor
	imprime(vivos, pos)
}

func imprime(vivos []int, espada int) {
	fmt.Print("[ ")
	for i := 0; i < len(vivos); i++ {
		if i == espada {
			fmt.Printf("%d> ", vivos[i])
		} else {
			fmt.Printf("%d ", vivos[i])
		}
	}
	fmt.Println("]")
}
