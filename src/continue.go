package main

import "fmt"

func main() {
    sum := 0
    for num := 1; num <= 100; num++ {
        if num % 5 == 0 {
            continue
        }
        sum += num
    }
    fmt.Println("sum 1..100 excluding nums divisible by 5 is", sum)
}
