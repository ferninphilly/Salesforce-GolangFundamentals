package main

import "fmt"

// START OMIT
type MyError struct {
    Num int
    Msg string
}

func (e MyError) Error() string {
    return fmt.Sprintf("%s (%d)", e.Msg, e.Num)
}

func do_something() (int, error) {
    return 0, MyError{42, "failure message"}
}

func main() {
    _, err := do_something()
    if me, ok := err.(MyError); ok { // assert err is of type MyError // HL
        fmt.Println("Num =", me.Num)
        fmt.Println("What =", me.Msg)
    }
}
// END OMIT
