package main

import "fmt"

func main() {
    var num int
    fmt.Print("Start the countdown at what number? ")
    fmt.Scanln(&num) 

    // above, doesn't work in playground, so...
    num = 10
    for ; num > 0; num-- { // HL
        fmt.Println(num)
    } // HL
    fmt.Println("Blast off!")
}
