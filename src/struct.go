package main

import "fmt"

// START OMIT
type mountain struct { // HL
    name      string // HL
    elevation int // HL
} // HL

func main() {
    k2 := mountain{"K2", 8611} // create a new struct
    fmt.Println(k2)
    fmt.Printf("%#v\n", k2)
}
// END OMIT
