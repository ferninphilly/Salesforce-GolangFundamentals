package main

import "fmt"

// START OMIT
import "sort"

func main() {
    strs := []string{"pear", "apple", "lemon"}
    sort.Strings(strs) // HL
    fmt.Println("sorted:", strs)
    // Sorting integers
    ints := []int{7, 2, -4}
    sort.Ints(ints) // HL
    fmt.Println("sorted:", ints)
    // Check if a slice is already sorted
    s := sort.StringsAreSorted(strs) // HL
    fmt.Println("Are strings sorted?", s)
}
// END OMIT
