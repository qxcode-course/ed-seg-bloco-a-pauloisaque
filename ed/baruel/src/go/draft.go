package main

import "fmt"

func main() {
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

	for i := 0; i < len(deveriaTer); i++ {
		for j := 0; j < len(figurinhas); j++ {
			if figurinhas[j] == deveriaTer[i].deviaTer && deveriaTer[i].slot == 1 {
				figurinhasRepetidas = append(figurinhasRepetidas, figurinhas[j])
			}
			if figurinhas[j] == deveriaTer[i].deviaTer && deveriaTer[i].slot == 0 {
				deveriaTer[i].slot = 1
			}
		}
	}

	var figurinhasFaltando []int = make([]int, 0)
	for i := 0; i < len(deveriaTer); i++ {
		if deveriaTer[i].slot == 0 {
			figurinhasFaltando = append(figurinhasFaltando, deveriaTer[i].deviaTer)
		}
	}

	if len(figurinhasRepetidas) == 0 {
		fmt.Println("N")
	} else {
		printSliceInts(figurinhasRepetidas)
	}

	if len(figurinhasFaltando) == 0 {
		fmt.Println("N")
	} else {
		printSliceInts(figurinhasFaltando)
	}

}

func printSliceInts(a []int) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
