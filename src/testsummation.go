package main

import (
	"fmt"
	"summation"
)

func main() {
	x := summation.Ex
	y := summation.Why
	z := summation.SumMeUp(x, y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%d + %d = %d\n", x, y, z)
}
