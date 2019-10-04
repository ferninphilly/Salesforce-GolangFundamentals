package main

// START OMIT
const (
	StatusOK              = iota   // 0
	StatusConnectionReset          // 1
	StatusOtherError               // 2
)

func main() {
	println(StatusOK)
	println(StatusConnectionReset)
}
// END OMIT
