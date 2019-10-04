package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

func sleeper(i int) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Int31n(20)) * time.Second)
	fmt.Println("Goroutine", i, "done!")
}

func main() {
	for i := 1; i <= 25; i++ {
		wg.Add(1)
        	go sleeper(i)
    	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
