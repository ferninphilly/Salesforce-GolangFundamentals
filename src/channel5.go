package main

import "fmt"
import "time"

// START OMIT
func worker(done chan bool) {
    for i := 1; i <= 5; i++ {
        fmt.Print("working...")
        time.Sleep(600 * time.Millisecond)
    }
    done <- true // we are done! // HL
}
func main() {
    done := make(chan bool, 1) // HL
    go worker(done)
    // block until we receive a notification
    <-done // HL
    fmt.Println("done!")
}
// END OMIT
