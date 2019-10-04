package main

import (
        "os"
        "fmt"
)

func main() {
    fmt.Println("All command-line args:", os.Args)
    fmt.Println("Number of args =", len(os.Args))
    fmt.Println("Arg 0 is", os.Args[0])
    fmt.Println("Args 1-3 are", os.Args[1:4])
}
