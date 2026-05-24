package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func embua(pen *Pen, dist float64) {
	if dist < 1 {
		return
	}
	pen.SetLineWidth(dist / 30)
	pen.Walk(dist)
	pen.Right(90)
	dist *= 0.95
	embua(pen, dist)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetRGB(0, 0, 250)
	pen.SetPosition(250, 250)
	pen.SetHeading(0)
	pen.SetPosition(100, 100)
	embua(pen, 300)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
