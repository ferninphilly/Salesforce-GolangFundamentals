package main

import "fmt"

// START OMIT
func check_type(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    check_type(21)
    check_type("hello")
    check_type(true)
}
// END OMIT
