package main

import "fmt"
import "sync"

var wait sync.WaitGroup
var count int = 0

func main() {
    wait.Add(2)
    go routine()
    go routine()
    wait.Wait()
    fmt.Println("count =", count)
}

func routine() {
    for i := 0; i < 2; i++ {
        value := count
        value++
        count = value
    }
    wait.Done()
}
