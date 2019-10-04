package main

// START OMIT
import "fmt"

func main() {
    var f float64 = 38730204.3832
    fmt.Printf("%v\n", f)
    fmt.Printf("%f\n", f)
    fmt.Printf("%.2f\n", f)
    var i int = 13
    fmt.Printf("|%5d|\n", i)
    fmt.Printf("|%-5d|\n", i)
    var b bool = false
    fmt.Printf("%t\n", b)
    fmt.Printf("%T %T %T\n", f, i, b)
}
// END OMIT
