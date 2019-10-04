package main

import "fmt"

// START OMIT
type MyInterface interface {
	Method()
}
type Text struct {
	String string
}

func (t *Text) Method() {
	if t == nil { // HL
		fmt.Println("<nil>") // HL
	} else {
		fmt.Println(t.String)
	}
}

func describe(i MyInterface) {
	fmt.Printf("(%v, %T)\n", i, i)
}
// END OMIT

// START2 OMIT
func main() {
	var i MyInterface
	var t *Text // HL
	
	i = t // HL
	describe(i) // HL
	i.Method()

	i = &Text{"hello"}
	describe(i)
	i.Method()
}
// END2 OMIT
