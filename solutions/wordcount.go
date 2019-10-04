package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
)

func main() {
	// Get filename from command-line args
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])

	// Bail out if file cannot be opened
	if err != nil {
		panic(err)
	}

	// Create a new sortedMap to hold the words and counts
	counts := map[string]int{}

	// Read a line at a time
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// Get rid of extra spaces by converting 2 spaces to 1
		line = strings.Replace(line, "  ", " ", -1)
		// Make it all lower case
		line = strings.ToLower(line)
		// Regexp to describe all non-alphanumeric characters
		reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
		if err != nil {
			panic(err)
		}
		// Remove all non-alphaumeric characters
		line = reg.ReplaceAllString(line, "")
		// Split line into words
		words := strings.Split(line, " ")

		// For each word, update it's count in the map
		for i := range words {
			counts[words[i]]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Iterate through map, printing out key and value
	for key, val := range counts {
		fmt.Println(key, val)
	}
}
