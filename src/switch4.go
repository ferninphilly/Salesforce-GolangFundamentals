package main

// START OMIT
import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
    // get current time, extract seconds (and discard hours and minutes)
    _, _, sec := time.Now().Clock() 
    // use seconds to seed the random number generator
    rand.Seed(int64(sec))
    r := rand.Int()

    switch { // HL
    case r % 2 == 0:
        fmt.Println("Random number is odd")
    default:
        fmt.Println("Random number is even")
    } // HL
}
// END OMIT
