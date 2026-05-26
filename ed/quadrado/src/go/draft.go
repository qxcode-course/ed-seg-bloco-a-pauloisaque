package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	square(n)
}

func square(n int) int {
	if n == 1 {
		fmt.Println("1^2 = 1")
		return 1
	}
	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)
	prev := square(n - 1)
	res := prev + 2*(n-1) + 1
	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, res)
	return res
}
