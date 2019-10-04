package main

import "fmt"
import "time"

func main() {
// START OMIT
	jobs := make(chan int, 5) // holds work to be done
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // more will be false if channel closed // HL
			if more {
				fmt.Println("running job", j)
				time.Sleep(time.Second)
			} else {
				fmt.Println("no more work to do")
				done <- true // signal main program // HL
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ { // sends 3 jobs to the worker over the jobs channel, then closes it // HL
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // HL
	fmt.Println("sent all jobs")

	<-done // await worker // HL
// END OMIT
}
