package main
import "os"
import "fmt"
import "os/signal"
import "syscall"

func main() {
    sigs := make(chan os.Signal, 1) // channel to receive os.Signal notifications // HL
    done := make(chan bool, 1) // channel to notify us when the program can exit
    // Register the given channel to receive signal notifications
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // HL
    go func() { // Execute a blocking receive for signals // HL
        sig := <-sigs // HL
        fmt.Println("\n", sig)
        done <- true // HL
    }() // HL
    fmt.Println("awaiting signal") // Wait here until goroutine is done
    <-done // HL
    fmt.Println("exiting")
}
