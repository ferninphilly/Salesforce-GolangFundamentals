package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
    _, _, sec := time.Now().Clock()
    rand.Seed(int64(sec))
// START OMIT
    num := rand.Int() % 1111
    fmt.Println("random number is", num)

    switch {
    case num > 1000:
        fmt.Println("greater than 1000")
        fallthrough // HL

    case num > 100:
        fmt.Println("greater than 100")
        fallthrough // HL

    case num > 10:
        fmt.Println("greater than 10")
	
    default:
        fmt.Println("less than 10")
    }
// END OMIT
}
