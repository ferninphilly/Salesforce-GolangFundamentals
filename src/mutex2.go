package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"sync/atomic" 
)

var state = make(map[int]int) // For our example the `state` will be a map
var mutex = &sync.Mutex{}     // This `mutex` will synchronize access to `state`
var readOps uint64  // We'll keep track of how many read...
var writeOps uint64 // ...and write operations we do

func read_state() {
    total := 0
    for {
        key := rand.Intn(5)           // pick a random map entry
        mutex.Lock()                  // lock the shared state // HL
        total += state[key]           // read it!
        mutex.Unlock()                // unlock // HL
        atomic.AddUint64(&readOps, 1) // atomically increment count // HL
        time.Sleep(time.Millisecond)
    }
}
func write_state() {
    for {
        key := rand.Intn(5)            // pick a random map position
        val := rand.Intn(100)          // pick a random number to write
        mutex.Lock()                   // lock the shared state // HL
        state[key] = val               // write it!
        mutex.Unlock()                 // unlock // HL
        atomic.AddUint64(&writeOps, 1) // atomically increment count
        time.Sleep(time.Millisecond)
    }
}

// START OMIT
func main() {
    for r := 0; r < 100; r++ { // Start 100 goroutines to execute repeated reads // HL
        go read_state()
    }
    for w := 0; w < 10; w++ {  // Start 10 goroutines to simulate writes // HL
        go write_state()
    }
    time.Sleep(time.Second)
    readOpsFinal := atomic.LoadUint64(&readOps) // Report final operation counts // HL
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps) // HL
    fmt.Println("writeOps:", writeOpsFinal)
    mutex.Lock() // HL
    fmt.Println("state:", state)
    mutex.Unlock() // HL
}
// END OMIT
