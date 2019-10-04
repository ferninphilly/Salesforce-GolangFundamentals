package main

import "fmt"

func address(title string, slice []string) {
	fmt.Printf("slice %s...%v\n", title, slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Address of %s[%d]: %p\n", title, i, &slice[i])
	}
}

func main() {
	s := make([]string, 3) // slice of strings of size 3
	s[0] = "zero"
	s[1] = "one"
	s[2] = "two"
	address("s", s)
	s2 := append(s, "four", "five", "six", "seven") // incorrect invocation of append()
	address("s2", s2)
	address("s", s)
	s = append(s, "eight")
	address("s", s)
	address("s2", s2)
	s2 = append(s2, "nine", "ten")
	address("s2", s2)
}
