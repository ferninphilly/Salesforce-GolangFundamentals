package main

import "fmt"

// START OMIT
func main() {
    f()
    fmt.Println("Back from f()")
}
func f() {
    defer func() {
        if r := recover(); r != nil { // HL
            fmt.Println("Recovered in f", r) // HL
        } // HL
    }()
    fmt.Println("Calling g()")
    g(0)
    fmt.Println("Back from g()")
}
func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
//END OMIT
