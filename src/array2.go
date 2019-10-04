package main

import "fmt"

// START OMIT
func main() {
    // ellipsis (...) can be used in place of size
    // if array is being intialized
    nums := [...]float32{ 1.4, 2.67, -3.56 } // HL
    fmt.Println("len(nums) =", len(nums))
    chars := [...]rune{ 'G', 'o', 'l', 'a', 'n', 'g', '€' } // HL
    fmt.Printf("%q\n", chars) // %q == quoted
    fmt.Println(chars) // values are actually int32
}
// END OMIT
