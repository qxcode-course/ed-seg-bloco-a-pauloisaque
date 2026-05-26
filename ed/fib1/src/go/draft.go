package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(rabbit(n, k))
}

func rabbit(n int, k int) int {
	if n <= 2 {
		return 1
	}
	return rabbit(n-1, k) + k*rabbit(n-2, k)
}
