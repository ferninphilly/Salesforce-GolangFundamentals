package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"sync/atomic"
)

var sharedint uint64
var sharedcount uint64
var wg sync.WaitGroup

func readwrite(n int) {
	defer wg.Done()
	iterations := rand.Int() % 1000 + 1000
	atomic.AddUint64(&sharedcount, uint64(iterations))
	//fmt.Printf("goroutine %d running for %d iterations\n", n, iterations)

	for i := 1; i <= iterations; i++ {
		atomic.AddUint64(&sharedint, 1)
		//sharedint++
		time.Sleep(time.Microsecond)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go readwrite(i)
	}
	wg.Wait()
	fmt.Println("value of shared int =", sharedint)
	fmt.Println("value of shared count =", sharedcount)
}

