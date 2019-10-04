package main

import "fmt"

// START OMIT
func main() {
    sum := 0
    for i := 1; i <= 100; i++ { // HL
        sum += i
    } // HL
    fmt.Println("sum of 1..100 is", sum)
}
// END OMIT
