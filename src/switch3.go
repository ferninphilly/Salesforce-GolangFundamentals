package main

import (
	"fmt"
	"time"
)

func main() {
    // time in playground is a constant, so this really
    // should be run on your machine...
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
}
