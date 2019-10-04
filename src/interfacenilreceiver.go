package main

import "fmt"

type MyInterface interface {
    Method()
}

type Text struct {
    String string
}

func describe(i MyInterface) {
    fmt.Printf("(%v, %T)\n", i, i)
}

// START OMIT
func (t *Text) Method() {
    if t == nil {
        fmt.Println("<nil>")
    } else {
        fmt.Println(t.String)
    }
}

func main() {
    var inter MyInterface
    inter.Method()
}
// END OMIT
