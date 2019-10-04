package main

import "fmt"

// START OMIT
type MyInterface interface { // HL
    Method() // for something to be of type MyInterface it must implement this method // HL
} // HL

type MyText string

func (text MyText) Method() { // this method means type MyText implements MyInterface // HL
    fmt.Println(text)
}

func main() {
    var inter MyInterface = MyText("Hyderabad")
    inter.Method()
}
// END OMIT
