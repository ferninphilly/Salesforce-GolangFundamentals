package main

import "fmt"
// START OMIT
import "math"
type triangle struct {
    side1, side2, side3 int
}

// compute area of triangle using Heron's formula
func (t triangle) area() int {
    s := (t.side1 + t.side2 + t.side3) / 2
    return int(math.Sqrt(float64(s * (s - t.side1) * (s - t.side2) * (s - t.side3))))
}

// compute perimeter of triangle
func (t triangle) perimeter() int {
    return t.side1 + t.side2 + t.side3
}

func main() {
    t := triangle{3, 4, 5}
    fmt.Println("area: ", t.area(), "\nperim:", t.perimeter()) // HL
}
// END OMIT
