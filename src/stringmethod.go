package main

import "fmt"
//START OMIT
import "strings"

type MyStr string // can't define a method on a basic type

func (s MyStr) Uppercase() string { // HL
    return strings.ToUpper(string(s)) 
}

func main() {
    s := MyStr("Golang is fun!")
    fmt.Println(s.Uppercase())
}
//END OMIT
