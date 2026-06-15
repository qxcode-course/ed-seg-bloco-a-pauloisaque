package main
import "fmt"
func main() {
    var a string
    fmt.Scan(&a)
    for i := 0; i < len(a); i++ {
        if a[i] == "(" {
            for j := i; j < len(a); j++ {
                if a[j] == ")" {
                    break;
                }
            }
        }
    }
}