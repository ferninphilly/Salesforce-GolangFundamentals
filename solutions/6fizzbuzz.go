// Write a function which accepts an integer and returns
//
// "fizz" if the number is divisible by 3
// "buzz" if the number is divisible by 5
// "fizzbuzz" if the number is divisible by BOTH 3 and 5
// otherwise it just returns the string version of the number, e.g., "4"
// 
// Test your function with these inputs: 3, 5, 15, 4
package main

import "fmt"

func fizzbuzz(n int) string {
	var fizz, buzz string

	if n % 3 == 0 {
		fizz = "fizz"
	}
	if n % 5 == 0 {
		buzz = "buzz"
	}
	if fizz + buzz == "" {
		return fmt.Sprintf("%d", n)
	}
	return fizz + buzz
}

func main() {
	testnums := [...]int{3, 5, 15, 4}

	for _, num := range testnums {
		fmt.Println(num, fizzbuzz(num))
	}
}
