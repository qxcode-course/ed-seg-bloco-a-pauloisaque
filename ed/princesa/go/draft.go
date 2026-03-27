package main

import "fmt"

type vivo struct {
    num int32
    taVivo bool
}

func main() {
	var pessoas int32
	var espada int32
	fmt.Scan(&pessoas, &espada)

	var vivos = make([]vivo, pessoas)
	for i := int32(0); i < pessoas; i++ {
		vivos[i].num = i + 1
        vivos[i].taVivo = true
	}

    fmt.Print(knext(vivos, espada))

}

func knext (vivos []vivo, espada int32) string {
    for i := range vivos {
        if i > len(vivos) {
            i = 0
        }
        if vivos[espada].taVivo == true {
            vivos[i + 1].taVivo = false
        } 
    }
    
    
}