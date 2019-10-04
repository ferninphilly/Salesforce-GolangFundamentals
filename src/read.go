package main

import "os"

// START OMIT
func check(e error) { // helper function for error checking
    if e != nil {
        panic(e)
    }
}

func main() {
    buf := make([]byte, 1024) // create a buffer  // HL
    f, e := os.Open("/etc/passwd") // HL
    check(e)
    defer f.Close() // idiomatic to defer a Close just after Open // HL
    for {
        n, e := f.Read(buf) // HL
        if n == 0 { // 0 means we've hit EOF
            break
        }
        check(e)
        os.Stdout.Write(buf[:n]) // HL
    }
}
// END OMIT
