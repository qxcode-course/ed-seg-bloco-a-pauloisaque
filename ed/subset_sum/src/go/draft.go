package main

import (
	"fmt"
)

func main() {
	var n, sum int
	fmt.Scan(&n, &sum)
	numeros := make([]int, n)
	path := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}
	result := false
	isSubset(path, numeros, sum, &result)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isSubset(path []int, numeros []int, sum int, result *bool) {
	if len(path) == 2 {
		if path[0] + path[1] == sum {
			*result = true
		}
		return
	}
	for i := 0; i < len(numeros); i++ {
		path = append(path, numeros[i])
		isSubset(path, numeros, sum, result)
		path = path[:len(path)-1]
	}
}