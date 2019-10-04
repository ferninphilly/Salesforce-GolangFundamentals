package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
var wg sync.WaitGroup // HL

func process(thing string) {
	defer wg.Done() // decrement counter when goroutine completes // HL
	time.Sleep(time.Duration(len(thing) * 100) * time.Millisecond)
	fmt.Println(thing, "processed")
}

func main() {
	var items = []string{
		"short",
		"longer",
		"even longer",
		"supercalifragilisticexpialidocious",
	}
	for _, item := range items {
		wg.Add(1) // increment the WaitGroup counter // HL
		go process(item) // HL
	}
	fmt.Println("main waiting...")
	wg.Wait() // wait for all processing to complete // HL
}
// END OMIT
