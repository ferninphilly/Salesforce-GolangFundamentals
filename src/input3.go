package main

import "fmt"

func main() {
    var word string
    var int1, int2 int
    var float1 float32

    fmt.Print("Enter text: ")
    fmt.Scanf("%d %f %d\n%s", &int1, &float1, &int2, &word)
    fmt.Println("I think you just entered...", int1, float1, int2, word)
}
