package main

import "fmt"

// START OMIT
func main() {
    // create a slice of strings of size 3
    s := make([]string, 3) // HL
    fmt.Println("initial slice:", s)
    s[0] = "zero" // set just like an array
    s[1] = "one"
    s[2] = "two"
    fmt.Println("now:", s)
    fmt.Println("get:", s[2]) // get like an array
    fmt.Println("len of slice:", len(s))
}
// END OMIT
