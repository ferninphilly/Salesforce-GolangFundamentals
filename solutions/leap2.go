package main

import "fmt"

func leap(year int) bool {
	switch {
	case year % 400 == 0:
		return true
	case year % 100 == 0:
		return false
	case year % 4 == 0:
		return true
	}
	return false
}

func main() {
	var year int

	fmt.Print("Enter a year: ")
 	fmt.Scanf("%d", &year)
	if ! leap(year) {
		fmt.Print("not a ")
	}
	fmt.Println("leap year")
}
