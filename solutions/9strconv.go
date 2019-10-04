// Combine the `strconv` exercise with an infinite loop so your code reads
// a line of input from the user, expecting it to be an integer.
// If you can't convert to an integer, print an error message and prompt
// the user again stop when the user complies.

package main

import "fmt"
import "strconv"

func getint() int {
	var str string

	for {
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&str)
		if res, err := strconv.Atoi(str); err != nil {
			fmt.Println("...not an integer, try again!")
		} else {
			return res
		}
	}
}

func main() {
	fmt.Println("You entered", getint())
}
