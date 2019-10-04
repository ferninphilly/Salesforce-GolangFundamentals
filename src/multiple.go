package main

import "fmt"

func addmul(x, y int) (int, int) {
    return x + y, x * y
}

func main() {
    s, p := addmul(3, 5)
    fmt.Println("sum =", s, "product =", p)
    _, r := addmul(4, -8)
    fmt.Println("product =", r)
}
