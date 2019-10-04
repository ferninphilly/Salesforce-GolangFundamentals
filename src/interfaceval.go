package main

import (
	"fmt"
	"math"
)

type MyInterface interface {
    Method()
}

type Text struct {
    String string
}

func (t *Text) Method() {
    fmt.Println(t.String)
}

type MyFloat float64

func (f MyFloat) Method() {
    fmt.Println(f)
}

// START OMIT
func describe(i MyInterface) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    var inter MyInterface
    inter = &Text{"Hello"} // pointer to Text object // HL
    describe(inter)
    inter.Method()
    inter = MyFloat(math.Pi) // MyFloat object // HL
    describe(inter)
    inter.Method()
}
// END OMIT
