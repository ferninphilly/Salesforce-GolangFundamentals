// Read in a string from the user
// Convert it to int using `strconv.Atoi()`
// Don't forget to check for errors!

package main

import "fmt"
import "strconv"

func main() {
	var str string

	fmt.Print("Enter an integer: ")
	fmt.Scanln(&str)
	if res, err := strconv.Atoi(str); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You entered", res)
	}
}
