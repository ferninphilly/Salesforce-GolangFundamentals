// Modify your square root function such that the loop terminates once the
// value has stopped changing (or only changes by a very small amount)
//
// See how many iterations it takes
//
// How close are your function's results to `math.Sqrt` in the standard
// library?

package main

import "fmt"
import "math"

func sqrt(x float64) float64 {
	currguess := 1.0
	prevguess := 1.1
	numguesses := 1

	for math.Abs(prevguess - currguess) > 1e-6 {
		prevguess = currguess
		currguess = prevguess - (prevguess * prevguess - x) / (2 * prevguess)
		fmt.Printf("Guess %d = %.10f\n", numguesses, currguess)
		numguesses++
	}
	return currguess
}

func main() {
	var num float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	res := sqrt(num)
	fmt.Printf("%.10f (%.10f)\n", res, math.Sqrt(num))
}
