// Write a program which repeatedly asks to enter a word
//
// Your program should quit when the user enters "quit"
//
// If the length of the word is even, your program should print the word
//
// If the length of the word is odd, your program should defer the
// printing of the word

package main

import "fmt"

func main() {
	var word string

infinite:
	for {
		fmt.Print("Enter a word: ")
		fmt.Scanln(&word)
		switch {
		case word == "quit":
			break infinite
		case len(word)%2 == 0:
			fmt.Println(word)
		default:
			defer fmt.Println(word)
		}
	}
}
