package main

import "fmt"

func main() {
loop: // HL
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            if i*j == 9 {
                fmt.Println("---")
                continue loop // skip rest of this loop // HL
            }
            fmt.Println(i, j)
        }
    }
}
