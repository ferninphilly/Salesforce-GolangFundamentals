package main

// START OMIT
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("poem.txt") // HL
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f) // HL
	for scanner.Scan() { // HL
		fmt.Println(scanner.Text()) // HL
	} // HL
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
// END OMIT
