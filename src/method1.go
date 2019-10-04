// compute area of triangle using Heron's formula
func (t triangle) area() int { // HL
    s := (t.side1 + t.side2 + t.side3) / 2
    return int(math.Sqrt(float64(s * (s - t.side1) * (s - t.side2) * (s - t.side3))))
}
