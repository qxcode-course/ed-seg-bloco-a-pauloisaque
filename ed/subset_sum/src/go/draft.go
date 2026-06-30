package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n, sum int
	fmt.Scan(&n, &sum)
	numeros := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}
	var isSubsetSum func(*Queue[int])
	path := NewQueue[int]()

	isSubsetSum = func(path *Queue[int]) {
		if path.items.Len() == 2 {
			if path.items.Front().Value.(int) + path.items.Front().Next().Value.(int) == sum {
				fmt.Println("true")
			}
			return
		}
		for i := 0; i < len(numeros); i++ {
			number := numeros[i]
			path.Enqueue(number)
			isSubsetSum(path)
			path.Dequeue()
		}
	}
	isSubsetSum(path)
}

type Queue[T any] struct {
	items *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: list.New()}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items.PushBack(item)
}

func (q *Queue[T]) Dequeue() T {
	if q.items.Len() == 0 {
		var zero T
		return zero // Return zero value if queue is empty
	}
	front := q.items.Front().Value.(T)
	q.items.Remove(q.items.Front())
	return front
}

func (q *Queue[T]) IsEmpty() bool {
	return q.items.Len() == 0
}

func (q *Queue[T]) String() string {
	var result string
	for e := q.items.Front(); e != nil; e = e.Next() {
		result += fmt.Sprint(e.Value.(T)) + " "
	}
	return "[ " + result + "]"
}
