package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(blocks(n))
}

func blocks(n int) int {
	if n == 1 {
		return 20
	}
	return blocks(n-1) + 8
}
