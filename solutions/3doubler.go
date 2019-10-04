// Write a function `doubler` which takes an integer
// argument and doubles it
package main

import "fmt"

func doubler(val *int) {
	*val *= 2
}

func main() {
	x := 5
	doubler(&x)
	fmt.Println(x)
}

