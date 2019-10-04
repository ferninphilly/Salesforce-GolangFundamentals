package main

import "fmt"

// if you name the parameters, you don't have to return them
// directly in the return statement...but I don't recommend it
func multival(x int) (square, cube int) { // HL
    square = x * x // HL
    cube = x * x * x // HL
    return
}

func main() {
    fmt.Println(multival(3))
}
