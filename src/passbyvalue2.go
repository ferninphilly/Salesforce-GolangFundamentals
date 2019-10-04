package main

import "fmt"

// the argument to this function is a pointer to an integer
// "star int" = pointer to int
func increment(nump *int) { // HL
    fmt.Printf("the value of nump is %v\n", nump)
    // now we increment the object that nump is pointing to
    *nump++ // HL
}

func main() {
    val := 1
    // this time we call increment, but pass a pointer to the val variable
    ptr_to_val := &val // HL
    fmt.Printf("ptr_to_val is the address of val... %v\n", ptr_to_val)
    increment(ptr_to_val)
    fmt.Println("Value after incrementing:", val)
}
