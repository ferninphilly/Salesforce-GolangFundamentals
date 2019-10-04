import "sync/atomic" // also "math/rand" and "time"

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
