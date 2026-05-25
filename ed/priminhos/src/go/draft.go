package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	primos := listaPrimos(n, 2, []int{})
	imprime(primos, 0)
}

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

func listaPrimos(n, atual int, primos []int) []int {
	if len(primos) == n {
		return primos
	}
	if ehPrimo(atual) {
		primos = append(primos, atual)
	}
	return listaPrimos(n, atual+1, primos)
}

func imprime(primos []int, i int) {
	if i == len(primos) {
		fmt.Println("]")
		return
	}
	if i == 0 {
		fmt.Print("[")
	} else {
		fmt.Print(", ")
	}
	fmt.Print(primos[i])
	imprime(primos, i+1)
}
