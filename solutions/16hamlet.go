package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Sorting a map by values is not nearly as straightforward
// in Go as it is in Python. The code below is copied from
// https://groups.google.com/forum/#!topic/golang-nuts/FT7cjmcL7gw
type sortedMap struct {
	m map[string]int
	s []string
}

// Return length of underlying map
func (sm *sortedMap) Len() int {
	return len(sm.m)
}

// "Less than" comparison compares values not keys
func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

// Swap values not keys
func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// Return a slice of sorted stringsa representing the
// keys, ordered by their corresponding values
func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

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
	counts := sortedMap{map[string]int{}, []string{}}

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
			counts.m[words[i]]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Get keys sorted by counts (values)
	keys := sortedKeys(counts.m)

	// Iterate through map, printing out key and value
	for _, key := range keys {
		fmt.Println(key, counts.m[key])
	}
}
