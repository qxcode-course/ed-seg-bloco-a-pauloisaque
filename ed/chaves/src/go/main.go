package main

func main() {
	lista := NewQueue[T]()
	times := "ABCDEFGHI"
	for i := 0; i < len(times); i++ {
		lista.Enqueue(times[i])
	}
}
