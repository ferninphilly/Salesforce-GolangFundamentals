package main

import "fmt"

// START OMIT
func product(nums ...int) {
    prod := 1
    for _, num := range nums {
        prod *= num
    }
    fmt.Println(nums, prod)
}

func main() {
    // Variadic functions can be called in the usual way with individual arguments
    product(3, 4)
    product(2, 3, 4, 5, 6)

    // If you already have multiple args in a slice, (which we haven't seen yet)
    // apply them to a variadic function using func(slice...) like this
    nums := []int{2, 3, 4, -1}
    product(nums...)
}
// END OMIT
