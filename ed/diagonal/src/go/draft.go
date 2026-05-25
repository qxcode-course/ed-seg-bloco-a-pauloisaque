package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	diagonal(s, 0)
}

func diagonal(s string, k int) {
	if len(s) == 0 {
		return
	}
	spaces(k)
	fmt.Println(string(s[0]))
	diagonal(s[1:], k+1)
}

func spaces(k int) {
	if k == 0 {
		return
	}
	fmt.Print(" ")
	spaces(k - 1)
}
