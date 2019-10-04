package main

import "fmt"

// this function takes three arguments: 1) a function which takes two ints
// as parameters, 2) an int, and 3) another int and then invokes that function,
// passing the two ints as parameters to the called function
func funcrunner(f func(x, y int), x, y int) { // HL
    f(x, y)
}

// this function takes 2 ints and prints their values
func printer(a, b int) {
    fmt.Println("a and b are", a, b)
}

func main() {
    // pass the printer function as the first argument
    // and 1 and 3 as the additional arguments
    funcrunner(printer, 1, 3) // HL
}
