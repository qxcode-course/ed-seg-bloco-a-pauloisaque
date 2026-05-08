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

	fmt.Println(verificarRepetidas(slotsAlbum, quantasTem, figurinhas))
}

func verificarRepetidas(slotsAlbum int, quantasTem int, figurinhas []int) {
	var deveriaTer []int
	for i := 0; i < slotsAlbum; i++ {
		deveriaTer[i] = i + 1
	}

}

func verificaRepetidas(slotsAlbum int, quantasTem int, figurinhas []int) []int {

}

func verificaFaltantes(slotsAlbum int, quantasTem int, figurinhas []int) []int {

}
