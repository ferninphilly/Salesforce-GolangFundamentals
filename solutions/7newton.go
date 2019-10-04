// Implement a square root function using Newton's method:
//
// - starting with some guess for the square root of x, we can adjust it
// based on how close guess² is to x, producing a better guess:

// guess = guess − (guess² − x) / (2 * guess)

// repeating the above makes the guess better and better

// use a starting guess of 1.0, regardless of the input (it works quite well)

//  repeat the calculation 10 times and print each guess along the way

package main

import "fmt"

func sqrt(x float64) float64 {
	currguess := 1.0
	prevguess := 1.1

	for count := 1; count <= 10; count++ {
		prevguess = currguess
		currguess = prevguess - (prevguess * prevguess - x) / (2 * prevguess)
		fmt.Println("Guess =", currguess)
	}
	return currguess
}

func main() {
	var num float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Println(sqrt(num))
}
