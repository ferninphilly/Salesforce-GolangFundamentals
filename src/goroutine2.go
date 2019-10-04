package main

import "fmt"

var done bool = false

func ask() {
        var input string
        fmt.Print("Enter some text: ")
        fmt.Scanln(&input)
        done = true
}

func main() {
        var square, count uint64 = 0, 0
        go ask() // HL
        for ; !done; count++ {
                square = count * count
        }
        fmt.Printf("While waiting, computed %d * %d = %d\n", count, count, square)
}
