package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func get_size(grid [][]rune) (int, int) {
	return len(grid), len(grid[0])
}

func is_value(grid [][]rune, pos Pos, value rune) bool {
	nl, nc := get_size(grid)
	l := pos.l
	c := pos.c
	if c < 0 || c >= nc || l < 0 || l >= nl {
		return false
	}
	return grid[l][c] == value
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	pos := Pos{l, c}
	stack.Push(pos)
	visited := make(map[Pos]bool)
	visited[pos] = true
	for !stack.IsEmpty() {
		elem := stack.Pop()
		if !is_value(grid, elem, '#') {
			continue
		}
		grid[elem.l][elem.c] = 'o'

		for _, neib := range getNeig(elem) {
			if !visited[neib] {
				stack.Push(neib)
				visited[neib] = true
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
