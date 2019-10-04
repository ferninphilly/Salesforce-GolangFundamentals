package main

import "fmt"

// START OMIT
func main() {
    var z []int
    fmt.Println(z, len(z))
    if z == nil {
        fmt.Println("nil!")
    }
}
// END OMIT
