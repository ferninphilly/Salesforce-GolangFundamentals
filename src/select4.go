package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)
    // Multi-way non-blocking select // HL
    select { // HL
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default: // HL
        fmt.Println("no activity")
    } // HL
}
