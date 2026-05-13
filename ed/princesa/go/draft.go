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

	printar := func(v []int) {
		fmt.Print("[")
		for i := 0; i < len(v); i++ {
			if v[i] == 1 {
				if i == e {
					fmt.Printf(" %d>", i+1)
				} else {
					fmt.Printf(" %d", i+1)
				}
			}
		}
		fmt.Println(" ]")
	}

	var quantosJaForam int
	quantosJaForam = 0
	printar(vetor)
	for quantosJaForam < len(vetor)-1 {
		var proximoTrue int
		proximoTrue = proximo()
		vetor[proximoTrue] = 0
		quantosJaForam++
		e = proximo()
		printar(vetor)
	}
}
