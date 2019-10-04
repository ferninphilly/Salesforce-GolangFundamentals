package main

import "fmt"
import "time"

// START OMIT
func pinger(c chan string) {
    fmt.Println("pinger about to send")
    c <- "ping...ping...ping" // send a value into a channel ... blocking operation // HL
}

func main() {
    mychan := make(chan string) // Create a new channel
    go pinger(mychan) // run pinger() as a goroutine // HL
    fmt.Println("goroutine started...waiting 3 seconds before receiving")
    time.Sleep(3 * time.Second)
    msg := <-mychan // receive value from channel ... this is a blocking operation // HL
    fmt.Println(msg)
}
// END OMIT
