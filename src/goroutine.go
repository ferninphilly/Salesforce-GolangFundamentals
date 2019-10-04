package main

import (
    "fmt"
    "time"
)
func thrice(from string) {
    for i := 1; i <= 3; i++ {
        time.Sleep(10 * time.Millisecond)
        fmt.Println(from, i)
    }
}
func main() {
    thrice("direct") // no concurrency at this point
    go thrice("goroutine") // invoke the function in a goroutine // HL
    go thrice("hey!") // and another...concurrent with main program // HL
    time.Sleep(1000 * time.Millisecond) // sleep 1 second
}
