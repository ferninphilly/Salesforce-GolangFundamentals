package main

import "fmt"

func main() {
loop: // HL
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            if i*j == 9 {
                fmt.Println("---")
                break loop // skip rest of both loops // HL
            }
            fmt.Println(i, j)
        }
    }
}
