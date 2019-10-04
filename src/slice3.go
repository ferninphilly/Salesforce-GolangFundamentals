package main

import "fmt"

// START OMIT
func main() {
    slice := make([]string, 0)
    slice = append(slice, "apple", "lemon", "mango", "banana", "cherry", "lime")
    fmt.Println(slice)
    // Slices support a "slice" operator with the syntax slice[low:high]
    slice1 := slice[2:5] // HL
    fmt.Println("slice1:", slice1)
    // This slices up to (but excluding) `s[5]`.
    slice2 := slice[:5] // HL
    fmt.Println("slice2:", slice2)
    // And this slices up from (and including) `s[2]`.
    slice3 := slice[2:] // HL
    fmt.Println("slice3:", slice3)
    // We can declare and initialize a variable for slice
    // in a single line as well.
    t := []string{"raspberry", "orange", "guava"}
    fmt.Println("declared slice:", t)
}
// END OMIT
