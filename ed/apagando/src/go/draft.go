package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	sairam := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&sairam[i])
	}

	subtrairVetor(fila, sairam)

}

func subtrairVetor(fila []int, sairam []int) {
	for i := 0; i < len(fila); i++ {
		for j := 0; j < len(sairam); j++ {
			if fila[i] == sairam[j] {
				fila[i] = 0
			}
		}
	}

	novaFila := make([]int, len(fila)-len(sairam))
	for i := 0; i < len(novaFila); i++ {
		for j := 0; j < len(fila); j++ {
			if fila[j] != 0 {
				novaFila[i] = fila[j]
				fila[j] = 0
				break
			}
		}
	}

}
