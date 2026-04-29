package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2, lado3 float64
	fmt.Scan(&lado1, &lado2, &lado3)

	p := (lado1 + lado2 + lado3) / 2

	area := math.Sqrt(p * (p - lado1) * (p - lado2) * (p - lado3))

	fmt.Printf("%.2f\n", area)
}
