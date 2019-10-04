package main

import "fmt"

// START OMIT
// return the square of a number
func squareit(val float64) float64 {
    return val * val
}

// return sum of two values
func sum(x int, y int) int { // equivalent: func sum(x, y int) int
    return x + y
}

func main() { // return nothing
    fmt.Printf("%f\n", squareit(4.5))
    fmt.Println(sum(13, -2))
}
// END OMIT
