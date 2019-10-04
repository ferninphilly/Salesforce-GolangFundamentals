package main

import "fmt"

// START OMIT
func main() {
    for index, char := range "Golang" { // HL
        // print each letter as a rune (int32) and character
        fmt.Printf("%d %d %c\n", index, char, char)
    }
}
// END OMIT
