package main

import "time"
import "fmt"

// START OMIT
func main() {
    c1 := make(chan string, 1)

    go func() { // sleep 2 seconds, then send
        time.Sleep(2 * time.Second)
        c1 <- "Secret message"
    }()

    select { // wait up to 1 second to get message on c1 // HL
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second): // HL
        fmt.Println("Timed out; Message not received.")
    }
}
// END OMIT
