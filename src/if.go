package main

import "fmt"

func main() {
    x := 27
    if x % 2 == 0 {
        fmt.Println(x, "is even")
    } else {
        fmt.Println(x, "is odd")
    }
}
