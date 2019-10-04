package main

import "fmt"
// START OMIT
import "errors"

var CountErr = errors.New("Invalid count") // HL
var EmptyErr = errors.New("Can't star empty string!") // HL
// END OMIT

// Like Python * for strings
func stringstar(str string, count int) (string, error) {
    result := ""
    if count < 1 {
        return "", CountErr // HL
    }
    if str == "" {
        return "", EmptyErr // HL
    }
    for i := 1; i <= count; i++ {
        result += str
    }
    return result, nil // HL
}

func main() {
	_, err := stringstar("hello", 0) // HL
	fmt.Println("error:", err)
}
// END2 OMIT
