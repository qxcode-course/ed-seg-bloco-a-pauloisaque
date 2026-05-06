package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var slotsAlbum int
	var quantasTem int
	figurinhas := make([]int, quantasTem)
	fmt.Scan(&slotsAlbum)
	fmt.Scan(&quantasTem)
	for i := 0; i < quantasTem; i++ {
		fmt.Scan(&figurinhas[i])
	}

	var quaisRepetem = verificaRepetidas(slotsAlbum, quantasTem, figurinhas)
	var quaisFaltam = verificaFaltantes(slotsAlbum, quantasTem, figurinhas)

	fmt.Println()
}

func verificaRepetidas(slotsAlbum int, quantasTem int, figurinhas []int) []int {

}

func verificaFaltantes(slotsAlbum int, quantasTem int, figurinhas []int) []int {

}
