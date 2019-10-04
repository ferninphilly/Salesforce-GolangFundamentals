package main

import "fmt"

// START OMIT
func main() {
	messages := make(chan string)

	select { // non-blocking receive // HL
	case msg := <-messages:
		fmt.Println("received message", msg)
	default: // HL
		fmt.Println("no message received")
	}

	msg := "hi"
	select { // non-blocking send // HL
	case messages <- msg:
		fmt.Println("sent message", msg)
	default: // HL
		fmt.Println("no message sent")
	}
}
// END OMIT
