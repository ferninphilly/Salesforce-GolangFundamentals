package main

import "io/ioutil"
import "os"
import "fmt"

func check(e error) { // helper function for error checking
    if e != nil {
        panic(e)
    }
}

// START OMIT
func main() {
    // Dump a string (or just bytes) into a file
    d1 := []byte("hello\ngo\n") // HL
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644) // HL
    check(err)
    // For more granular writes, open a file for writing
    f, err := os.Create("/tmp/dat2") // HL
    check(err)
    defer f.Close() // HL
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2) // you can Write byte slices as you'd expect // HL
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
    n3, err := f.WriteString("writes\n") // write a string // HL
    fmt.Printf("wrote %d bytes\n", n3)
    f.Sync() // flushes writes to stable storage // HL
}
// END OMIT
