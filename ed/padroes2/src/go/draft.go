package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(codigo(n))
}

func codigo(n int) int {
	if n == 0 {
		return 0
	}
	return codigo(n-1) + 2*n + 1
}
