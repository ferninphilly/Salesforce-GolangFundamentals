package main

import "fmt"

// START OMIT
func main() {
    var a [10]int // array of 10 integers, initialized to 0 by default
    fmt.Println(a)
    for i := 0; i < len(a); i++ {
        a[i] = i + 101
    }
    fmt.Println(a)
    fmt.Println("a[3] =", a[3]) // array indexing
    // Arrays can be initialized when they are declared
    cities := [5]string{"Hyderabad", "Delhi", "Pune", "Kochi", "Chennai"}
    fmt.Println("Cities:", cities)
}
// END OMIT
