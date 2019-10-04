package main

// This function takes a function as a parameter. The function
// passed in must take an int argument and return an int.
func runsomething(f func(int) int, i int) int {
    return f(i)
}

func main() {
    // Here we call the above function and pass in an
    // anonymous function (or function literal)
    x := runsomething(
        func(a int) int { // HL
            return a * 2 // HL
        }, -5) // HL
    println(x)
}
