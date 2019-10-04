package main

import "fmt"

func fibonacci(c chan int, quit chan bool) {
    curr, next := 1, 1
    for {
	// INFINITE LOOP: Check if either
	// 1) we can send a Fibonacci number on
	//    the unbuffered channel
	// ...OR...
	// 2) we can read from the quit channel
	//    and if so, we quit (return)
	select {
	case c <- curr:
	    curr, next = next, curr+next

	case <-quit:
	    fmt.Println("main said stop!")
	    return
	}
    }
}

func main() {
    c := make(chan int) // Fibonacci number channel
    quit := make(chan bool) // quit channel

    go fibonacci(c, quit)
    for i := 1; i <= 10; i++ {
        num := <-c
	fmt.Println(num)
    }
    quit <- true
}
