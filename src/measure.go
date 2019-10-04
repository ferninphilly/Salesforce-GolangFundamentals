package main

import (
	"fmt"
	"math"
)

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// START OMIT
type geometry interface {
    area() float64
    perim() float64
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    // The circle and rect struct types both implement // HL
    // the geometry interface so we can use instances of // HL
    // of these structs as arguments to measure() // HL
    measure(r) // HL
    measure(c) // HL
}
// END OMIT
