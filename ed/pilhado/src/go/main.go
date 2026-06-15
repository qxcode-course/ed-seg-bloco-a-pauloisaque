package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, pos, endPos Pos, visited map[Pos]bool) bool {
	if !match(grid, pos, ' ') || visited[pos] {
		return false
	}
	visited[pos] = true
	if pos == endPos {
		grid[pos.l][pos.c] = '.'
		return true
	}
	for _, elem := range getNeig(pos) {
		if search(grid, elem, endPos, visited) {
			grid[pos.l][pos.c] = '.'
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	visited := make(map[Pos]bool, 0)
	search(grid, startPos, endPos, visited)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
