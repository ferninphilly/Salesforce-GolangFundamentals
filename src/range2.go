package main

import "fmt"

// START OMIT
func main() {
    nums := []int{5, 13, 27, 19, -1}
    sum := 0
    for _, num := range nums { // HL
        sum += num
    }
    fmt.Println("sum of", nums, "is", sum)
}
// END OMIT
