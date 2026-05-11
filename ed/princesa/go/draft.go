package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)
	e = e - 1
	josephusProblem(n, e)
}

func josephusProblem(n int, e int) {
	var vetor []int = make([]int, n)
	for i := 0; i < n; i++ {
		vetor[i] = 1
	}

	proximo := func() int {
		var res int
		for i := 0; i < len(vetor); i++ {
			valorAtual := (i + e) % len(vetor)
			valorVerificado := (valorAtual + 1) % len(vetor)
			if vetor[valorVerificado] == 1 {
				res = valorVerificado
				break
			}
		}
		return res
	}

	transformarProximoFalse := func() bool {
		var proximoTrue int
		proximoTrue = proximo()
		vetor[proximoTrue] = 0
		return false
	}

	var temUmElemento bool
	temUmElemento = false
	for temUmElemento == false {
		temUmElemento = transformarProximoFalse()
	}

	fmt.Println(proximo())

}

func transformaProximoEmFalse() {

}

func passarEspadaProximo() {

}
