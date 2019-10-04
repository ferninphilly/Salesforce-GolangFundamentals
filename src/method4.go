package main

import "fmt"
import "math"
type triangle struct {
    side1, side2, side3 int
}

// compute area of triangle using Heron's formula
// START OMIT
func (t *triangle) area() int { // HL
    s := (t.side1 + t.side2 + t.side3) / 2
    return int(math.Sqrt(float64(s * (s - t.side1) * (s - t.side2) * (s - t.side3))))
}

func (t *triangle) perimeter() int { // HL
    return t.side1 + t.side2 + t.side3
}

func (t *triangle) scale(factor int) { // HL
	t.side1 *= factor 
	t.side2 *= factor 
	t.side3 *= factor 
} 

func main() {
    t := triangle{3, 4, 5}
    t.scale(3)
    fmt.Println("sides: ", t.side1, t.side2, t.side3)
}
// END OMIT
