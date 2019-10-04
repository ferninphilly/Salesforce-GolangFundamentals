package main

import "fmt"
import "time"

func runner(baton chan int) {
    num := <-baton // wait to find out which runner I am
    fmt.Println("Runner", num, "running with the baton")

    if num != 4 {
	fmt.Println("Runner", num+1, "moves to the starting line")
        go runner(baton)
    }    
    time.Sleep(5 * time.Second) // this is me running
    if num == 4 {
	fmt.Println("Race over")
	return
    }
    fmt.Println("Runner", num, "done, passing baton to next runner")
    baton <- num+1
}

func main() {
    baton := make(chan int) // create channel
    go runner(baton) // start the first runner
    baton <- 1 // send "1" to runner 
    time.Sleep(30 * time.Second)
}
