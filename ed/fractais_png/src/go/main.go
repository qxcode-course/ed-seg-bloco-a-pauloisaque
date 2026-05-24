package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

//arvore.png

// func arvore(pen *Pen, dist float64) {
// 	if dist < 10 {
// 		return
// 	}
// 	ang := 35.0
// 	fator := 0.75
// 	pen.Walk(dist)
// 	//pen.SavePNG("tree.png")
// 	//var dummy rune
// 	//fmt.Scanf("%c", &dummy)
// 	pen.Right(ang)
// 	arvore(pen, dist*fator)
// 	pen.Right(-2 * ang)
// 	arvore(pen, dist*fator)
// 	pen.Right(ang)
// 	pen.Walk(-dist)
// }

// func main() {
// 	pen := NewPen(500, 500)
// 	pen.SetHeading(90)
// 	pen.SetPosition(250, 500)
// 	arvore(pen, 100)

// 	pen.SavePNG("arvore.png")
// 	fmt.Println("PNG file created successfully.")
// }

func circulos(pen *Pen, dist float64) {
	if dist < 50 {
		return
	}
	ang := 300.0
	fator := 0.75
	pen.DrawCircle(100)
	pen.SetHeading(90)
	//pen.SavePNG("tree.png")
	//var dummy rune
	//fmt.Scanf("%c", &dummy)
	pen.Right(ang)
	circulos(pen, dist*fator)
	pen.Right(-2 * ang)
	circulos(pen, dist*fator)
	pen.Right(ang)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetHeading(90)
	pen.SetPosition(250, 250)
	circulos(pen, 100)

	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}
