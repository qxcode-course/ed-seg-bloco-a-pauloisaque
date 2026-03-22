package main

import "fmt"

func main() {
	var helicoptero int32
	var policial int32
	var fugitivo int32
	var direcao int32
	fmt.Scan(&helicoptero)
	fmt.Scan(&policial)
	fmt.Scan(&fugitivo)
	fmt.Scan(&direcao)
	fmt.Print(consegueFugir(helicoptero, policial, fugitivo, direcao))
}

func consegueFugir(helicoptero int32, policial int32, fugitivoBuffer int32, direcao int32) string {
	var fugiu = ""
	for {
		if fugitivoBuffer == 16 && direcao == 1 {
			fugitivoBuffer = 0
		}
		if fugitivoBuffer == 0 && direcao == -1 {
			fugitivoBuffer = 15
		}
		if fugitivoBuffer == helicoptero {
			fugiu = "S"
			break
		}
		if fugitivoBuffer == policial {
			fugiu = "N"
			break
		}
		fugitivoBuffer += direcao
	}
	return fugiu + "\n"
}
