package main

import "fmt"

func main() {
    num := -1

    for num < 0 { // HL
        fmt.Print("Enter a positive number: ")
        fmt.Scanln(&num) // won't work in playground!
    } // HL

    fmt.Println("Thanks for following directions!")
}
