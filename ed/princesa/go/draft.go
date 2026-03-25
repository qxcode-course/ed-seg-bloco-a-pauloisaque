package main

import "fmt"

func main() {
	var pessoas int32
	var espada int32
	fmt.Scan(&pessoas, &espada)

	var vivos []int32
	for i := int32(0); i < pessoas; i++ {
		vivos = append(vivos, i+1)
	}

	for i := range vivos {
		fmt.Print(vivos[i])
	}

}
