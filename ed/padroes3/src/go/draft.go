package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(calcular(n, m))
}

func calcular(n int, m int) int {
	if m == 1 {
		return 1
	}
	return calcular(n, m-1) + (n-2)*(m-1) + 1
}
