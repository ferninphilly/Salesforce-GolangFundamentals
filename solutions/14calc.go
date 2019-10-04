// Create four calculator functions: add, sub, mul, div
//
// Each function should take 2 integers and return an integer result
// Create a map which maps the strings representing the operators to
// the _functions_, that is, "+" would be mapped to `add()`, "-" would
// be mapped to `sub()`...
//
// Have your program read in lines like "2 + 4" and have it determine
// the result by parsing the line and then using the operator ("+" in
// this case) to find the appropriate function to invoke.

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func div(x, y int) int {
	return x / y
}

func mul(x, y int) int {
	return x * y
}

var dispatch = map[string]func(int, int) int {
	"+": add,
	"-": sub,
	"/": div,
	"*": mul,
}

func main() {
	var num1, num2 int
	var op string

	fmt.Scanf("%d %s %d", &num1, &op, &num2)
	fmt.Println(dispatch[op](num1, num2))
}
