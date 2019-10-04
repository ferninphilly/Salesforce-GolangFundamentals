package main

import "fmt"

func leap(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
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
