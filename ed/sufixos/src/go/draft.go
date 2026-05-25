package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	printSufixos(s, len(s)-1)
}

func printSufixos(s string, i int) {
	if i < 0 {
		return
	}
	fmt.Println(s[i:])
	printSufixos(s, i-1)
}
