package main

// START OMIT
import "fmt"
import "time"

func main() {
    // make a buffered channel of up to 4 strings
    bchan := make(chan string, 4)
    // We can send 4 values without a corresponding receive
    bchan <- "1"   // will not block // HL
    bchan <- "2"   // will not block // HL
    bchan <- "3"   // will not block // HL
    bchan <- "go!" // will not block // HL
    fmt.Println("4 values sent...")
    time.Sleep(1 * time.Second)
    for i := 1; i <= 4; i++ { // Later we can receive these values as usual
        fmt.Println(<-bchan)
    }
}
// END OMIT
