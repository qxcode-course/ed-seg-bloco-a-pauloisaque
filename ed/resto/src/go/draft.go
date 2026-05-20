package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	restofunc(num)
}

func restofunc(num int) int {
	resto := num % 2
	num = num / 2

	if num > 0 {
		restofunc(num)
	}
	fmt.Printf("%d %d\n", num, resto)
	return 0
}
