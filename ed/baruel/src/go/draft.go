package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var slotsAlbum int
	var quantasTem int
	fmt.Scan(&slotsAlbum)
	fmt.Scan(&quantasTem)
	figurinhas := make([]int, quantasTem)
	for i := 0; i < quantasTem; i++ {
		fmt.Scan(&figurinhas[i])
	}

	verificarRepetidas(slotsAlbum, quantasTem, figurinhas)
}

func verificarRepetidas(slotsAlbum int, quantasTem int, figurinhas []int) {
	type slots struct {
		deviaTer int
		slot     int
	}
	var deveriaTer []slots = make([]slots, slotsAlbum)
	for i := 0; i < slotsAlbum; i++ {
		deveriaTer[i] = slots{i + 1, 0}
	}
	var figurinhasRepetidas []int = make([]int, 0)

	for i := 0; i < quantasTem; i++ {
		for i := 0; i < len(deveriaTer); i++ {
			if figurinhas[i] == deveriaTer[i].deviaTer && deveriaTer[i].slot == 0 {
				deveriaTer[i].slot = figurinhas[i]
			}
		}
	}

	for i := 0; i < quantasTem; i++ {
		if figurinhas[i] == deveriaTer[i].deviaTer && deveriaTer[i].slot != 0 {
			figurinhasRepetidas = append(figurinhasRepetidas, figurinhas[i])
		}
	}

	var quantasEstaoFaltando int
	for i := 0; i < len(deveriaTer); i++ {
		if deveriaTer[i].slot == 0 {
			quantasEstaoFaltando++
		}
	}

	fmt.Println(figurinhasRepetidas)
	fmt.Println(quantasEstaoFaltando)

}
