package main

import "fmt"
import "strconv"

func sumdigits(n int) string {
	s := strconv.Itoa(n)
	return s
}

func main() {
 	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)
	fmt.Println(sumdigits(num))
}
