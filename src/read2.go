package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"bufio"
)

func check(e error) { // helper function for error checking
    if e != nil {
        panic(e)
    }
}

// START OMIT
func main() {
// Slurping a file's entire contents into memory
    dat, err := ioutil.ReadFile("/tmp/dat1")  // HL
    check(err)
    fmt.Print(string(dat))
    // or...
    f, err := os.Open("/tmp/dat1") // HL
    check(err)
    defer f.Close() // HL
    // Read up to 5 bytes from beginning of file.
    // Note how many actually were read.
    b1 := make([]byte, 5) // HL
    n1, err := f.Read(b1) // HL
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    o2, err := f.Seek(6, 0) // seek to known location in file and read from there // HL
    check(err)
    b2 := make([]byte, 2) // HL
    n2, err := f.Read(b2) // HL
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
// END OMIT
// START2 OMIT
    // io package provides helpful functions for file reading. Reads
    // like the ones above can be more robustly implemented with ReadAtLeast.
    o3, err := f.Seek(6, 0) // HL
    check(err)
    b3 := make([]byte, 2) // HL
    n3, err := io.ReadAtLeast(f, b3, 2) // HL
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
    _, err = f.Seek(0, 0) // no built-in rewind, but Seek(0, 0) does the trick // HL
    check(err)
    // bufio package implements a buffered reader–useful both for efficiency
    // with many small reads and additional reading methods it provides
    r4 := bufio.NewReader(f) // HL
    b4, err := r4.Peek(5) // HL
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
}
// END2 OMIT
