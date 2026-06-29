package main

import "fmt"

func (q *Queue[T]) quemGanhou(a int, b int) {
	for q.items.Len() > 1 {
		e1 := q.Dequeue()
		e2 := q.Dequeue()
		if a > b {
			q.Enqueue(e1)
			break
		}
		if b > a {
			q.Enqueue(e2)
			break
		}
	}
}

func main() {
	var a, b int

	lista := NewQueue[string]()

	times := "ABCDEFGHIJKLMNOP"
	for i := 0; i < len(times); i++ {
		lista.Enqueue(string(times[i]))
	}

	for i := 0; i < 15; i++ {
		fmt.Scan(&a, &b)
		lista.quemGanhou(a, b)
	}

	fmt.Println(lista.Dequeue())

	
	
}
