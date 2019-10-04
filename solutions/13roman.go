// Use a map to translate Roman numerals into their Arabic equivalents.
// Load the map with Roman numerals M (1000), D (500), C (100), L (50),
// X (10), V (5), I (1).
// Read in a Roman numeral
// Print Arabic equivalent
// Try it with MCLX = 1000 + 100 + 50 + 10 = 1160
// EXTRA: if you have time, deal with the case where a smaller number
// precedes a larger number,
// e.g., XC = 100 - 10 = 90, or MCM = 1000 + (1000-100) = 1900

package main

import "fmt"

func roman(numeral string) int {
	roman_to_arabic := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	} 
	arabic_list := make([]int, len(numeral))

        for index, rune := range numeral {
	    if val, inmap := roman_to_arabic[rune]; !inmap {
		fmt.Printf("Invalid Roman digit: %c\n", rune)
		return 0
	    } else {
	    	arabic_list[index] = val
	    }
	}
	total := arabic_list[len(numeral)-1]
	for index := 0; index < len(numeral) - 1; index++ {
     		if arabic_list[index] < arabic_list[index + 1] {
        		arabic_list[index] = -arabic_list[index]
		}
		total += arabic_list[index]
	}
	return total
}

func main() {
    var s string

    fmt.Print("Enter Roman numeral: ")
    fmt.Scanln(&s)
    fmt.Println(roman(s))
}
