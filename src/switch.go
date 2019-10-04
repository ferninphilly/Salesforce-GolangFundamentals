package main

import "fmt"

func main() {
    i := 2
    fmt.Print("The value ", i, " is ")

    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
}
