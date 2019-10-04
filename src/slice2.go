package main

import "fmt"

// START OMIT
func main() {
    s := make([]string, 3) // slice of strings of size 3
    fmt.Println("initial slice:", s)
    s[0] = "zero" // set just like an array
    s[1] = "one"
    s[2] = "two"
    fmt.Println("now:", s)
    fmt.Println("get:", s[2]) // get like an array
    fmt.Println("len of slice:", len(s))
    s = append(s, "three") // HL
    s = append(s, "four", "five") // HL
    fmt.Println("after append:", s)
    copy_of_s := make([]string, len(s)) // HL
    copy(copy_of_s, s) // HL
    fmt.Println("copy:", copy_of_s)
}
// END OMIT
