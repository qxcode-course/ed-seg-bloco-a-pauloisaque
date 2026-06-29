package main

import "fmt"

type Table struct {
	parenteses int
	colchetes  int
}

func main() {
	var a string
	fmt.Scan(&a)
	var tabela Table
	for i := 0; i < len(a); i++ {
		if a[i] == '(' {
			tabela.parenteses++
		}
		if a[i] == ')' {
			tabela.parenteses--
		}
		if a[i] == '[' {
			tabela.colchetes++
		}
		if a[i] == ']' {
			tabela.colchetes--
		}
	}
	if tabela.colchetes == 0 && tabela.parenteses == 0 {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}
