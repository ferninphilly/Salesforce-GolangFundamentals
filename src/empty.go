package main

import "fmt"

// START OMIT
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    var i interface{}
    describe(i)
    i = 42
    describe(i)
    i = "hello"
    describe(i)
}
// END OMIT
