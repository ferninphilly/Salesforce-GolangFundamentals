package main

import "time"
import "fmt"

func main() {
// START OMIT
    c1, c2 := make(chan string), make(chan string)
    go func() { // wait 3 seconds, then send
        time.Sleep(3 * time.Second)
        c1 <- "one"
    }()
    go func() { // wait 1 seconds, then send
        time.Sleep(1 * time.Second)
        c2 <- "two"
    }()
    for i := 0; i < 2; i++ { // await both of these values simultaneously
        select { // HL
        case msg1 := <-c1: // HL
            fmt.Println("received", msg1)
        case msg2 := <-c2: // HL
            fmt.Println("received", msg2)
        } // HL
    }
// END OMIT
}
