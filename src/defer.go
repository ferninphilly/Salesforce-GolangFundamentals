package main

import "fmt"

// START OMIT
func main() {
    for i := 1; i <= 4; i++ {
        defer fmt.Println("deferred", -i) // HL
        fmt.Println("regular", i)
    }
}
// END OMIT
