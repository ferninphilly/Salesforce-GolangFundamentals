package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    // set up a logger to include the date and time in microseconds
    log.SetFlags(log.Ldate | log.Lmicroseconds)
    // set up a prefix which for every log message
    log.SetPrefix("MyApp ")
    // Start program?
    log.Println("Kaaryakram shuroo karen")
    // do some stuff...
    time.Sleep(103871 * time.Microsecond)
    log.Println("Database connection dropped.")
    // cf. with standard fmt.Println
    fmt.Println("Database connection dropped.")
}

