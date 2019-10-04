package main

import "fmt"

// START OMIT
// It would be a compile-time error to try to receive on this channel
func ping(schan chan<- string, msg string) { // HL
    fmt.Printf("ping sending %#v\n", msg)
    schan <- msg
}

// Accepts one channel for receives and a second for sends
func pong(rchan <-chan string, schan chan<- string) { // HL
    msg := <-rchan // receive a message
    fmt.Printf("pong received %#v\n", msg)
    schan <- msg
    fmt.Println("and sent it...")
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "secret message")
    pong(pings, pongs)
    fmt.Printf("main received %#v\n", <-pongs)
}
// END OMIT
