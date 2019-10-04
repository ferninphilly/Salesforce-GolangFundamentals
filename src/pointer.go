package main

import "fmt"

func main() {
	x := 1
	px := &x // The & operator takes the address of what follows it // HL
	fmt.Printf("The value of x is %d\n", x)
	fmt.Printf("The address of x is %v\n", px)
	fmt.Printf("The canonical form of px is %#v\n", px)
}
