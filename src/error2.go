package main
import "fmt"

import "time"

// START OMIT
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string { // HL
	return fmt.Sprintf("at %v, %s", e.When, e.What) // HL
} // HL

func run() error {
	return MyError{
		time.Now(), 
		"it didn't work", }
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
// END OMIT
