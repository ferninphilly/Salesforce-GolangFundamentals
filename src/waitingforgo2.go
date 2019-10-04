package main

import (
    "fmt"
    "net/http"
    "sync"
)

// START OMIT
func main() {
    var wg sync.WaitGroup // HL

    var urls = []string{
        "http://www.golang.org/", "http://www.google.com/", "http://www.zzz1983x.com/",
    }
    for _, url := range urls {
        wg.Add(1) // increment the WaitGroup counter // HL
        go func(url string) { // launch goroutine to fetch URL // HL
            defer wg.Done() // decrement counter when goroutine completes // HL
            if resp, err := http.Get(url); err == nil { // HL
                fmt.Println(url, resp.Status)
            } else {
                fmt.Println(url, err)
            }
        }(url) // HL
    }
    wg.Wait() // Wait for all HTTP fetches to complete // HL
}
// END OMIT
