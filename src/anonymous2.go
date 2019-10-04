package main

// This function takes a function as a parameter. The function
// passed in must take an int argument and return an int.
func runsomething(f func(int) int, i int) int {
    return f(i)
}

func main() {
    // Here we make the variable lambda refer to
    // the function, and then pass lambda into the
    // function above.
    lambda := func(a int) int { // HL
        return a * 2 // HL
    } // HL
    println(runsomething(lambda, -5))
}
