package main

import "fmt"

// START OMIT
type MyInterface interface {
    Method()
}

func describe(i MyInterface) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    var i MyInterface
    describe(i)
    i.Method()
}
// END OMIT
