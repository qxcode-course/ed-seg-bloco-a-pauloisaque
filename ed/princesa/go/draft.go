package main

import "fmt"

type vivo struct {
	num    int32
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

// 1> 2 3 4

func knext(vivos []vivo, espada int32) string {
	var resultado = ""
    var qtdVivos = len(vivos)

	resultado += "[ "
	for i := range vivos {
		if vivos[i].taVivo == true && vivos[i].num == espada {
			resultado += fmt.Sprintf("%d", vivos[i].num) + "> "
			continue
		}
        if vivos[i].taVivo == true {
			resultado += fmt.Sprintf("%d ", vivos[i].num)
			continue
		}
	}
	resultado += "]"

	for i := range vivos {
        if vivos[i].num == espada {
            if i < len(vivos) - 1 && vivos[i + 1].taVivo == true {
                vivos[i+1].taVivo = false
                if espada + 2 > int32(len(vivos)) {
					espada = 1
				}
                espada += 2
                qtdVivos--
                continue
            }
        }
	}

    if qtdVivos == 1 {
        return resultado
    }

	return resultado + knext(vivos, espada)
}
