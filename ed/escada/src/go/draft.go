package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(maneiras(n))
}

func maneiras(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	return maneiras(n-1) + maneiras(n-3)
}
