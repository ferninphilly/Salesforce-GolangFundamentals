package main

import "fmt"

func main() {
    // make a buffered channel of up to 2 strings
    bchan := make(chan string, 2)

    bchan <- "1" // HL
    bchan <- "2" // HL
    bchan <- "3" // HL

    // Later we can receive these values as usual
    for i := 1; i <= 3; i++ {
        fmt.Println(<-bchan)
    }
}
