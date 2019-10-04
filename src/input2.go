package main

import "fmt"

func main() {
    var input string
    var n int
    fmt.Print("Enter text: ")
    fmt.Scan(&input, &n)
    fmt.Println("I think you just entered...", input, n)
}
