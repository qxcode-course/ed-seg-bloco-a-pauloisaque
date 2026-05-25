package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	return ehPrimoRec(n, 2)
}

func ehPrimoRec(n, d int) bool {
	if d*d > n {
		return true
	}
	if n%d == 0 {
		return false
	}
	return ehPrimoRec(n, d+1)
}

func enesimoPrimo(n int) int {
	return buscaPrimo(n, 2, 0)
}

func buscaPrimo(alvo, atual, cont int) int {
	if ehPrimo(atual) {
		cont++
		if cont == alvo {
			return atual
		}
	}
	return buscaPrimo(alvo, atual+1, cont)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(enesimoPrimo(n))

}
