package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Printf("%.2f\n", heron(a, b, c))
}

func heron(a float64, b float64, c float64) float64 {
	p := float64(a+b+c) / 2
	var area float64
	area = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return area
}
