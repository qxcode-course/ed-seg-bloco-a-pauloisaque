package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)
	e = e - 1
	josephusProblem(n, e)
}

func josephusProblem(n int, e int) {
	var vetor []int = make([]int, n)
	for i := 0; i < n; i++ {
		vetor[i] = i + 1
	}
	fmt.Println(vetor)

	for i := 0; i < len(vetor); i++ {
		valor := (i + e) % len(vetor)
		if vetor[valor+1] != 0 {
			vetor[valor+1] = 0
			e = (e + 2) % len(vetor)
		} else {
			continue
		}
	}

}
