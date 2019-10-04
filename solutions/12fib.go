// Write a function `fib` that returns a slice containing the first n
// Fibonacci numbers
//
// The Fibonacci sequence is 1, 1, 2, 3, 5, 8...
//
// Each number is the sum of the previous Fibonacci numbers
//
// e.g., `fib(6)` should return an int slice of size 6 containing
// 1, 1, 2, 3, 5, 8

package main

import "fmt"

func fib(n int) []int {
    s := make([]int, n)

    s[0], s[1] = 1, 1

    for i := 2; i < n; i++ {
        s[i] = s[i-1] + s[i-2]
    }
    return s
}

func main() {
    var num int

    fmt.Print("Fib for how many? ")
    fmt.Scanln(&num)
    fmt.Println(fib(num))
}

