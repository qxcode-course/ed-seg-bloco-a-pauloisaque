package main

import "fmt"

func main() {
	var name string
	var age int32
	fmt.Scan(&name)
	fmt.Scan(&age)
	fmt.Println(name, "eh", verifyAge(age))
}

func verifyAge(a int32) string {
	if a < 12 {
		return "crianca"
	}
	if a < 18 {
		return "jovem"
	}
	if a < 65 {
		return "adulto"
	}
	if a < 1000 {
		return "idoso"
	}
	return "mumia"
}
