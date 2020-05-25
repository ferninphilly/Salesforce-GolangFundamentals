package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	for {
		fmt.Println("Please enter an integer!")
		fmt.Scan(&num)
		g, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("I SAID please enter an integer!")
			continue
		}
		fmt.Printf("Thank you for following directions! %d is a number!\n", g)
		break
	}
}
