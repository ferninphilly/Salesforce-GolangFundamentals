package main

var _ int = setup() // happens first

func init() { // happens second // HL
    println("init!")
}

func init() { // happens third // HL
    println("something else")
}

func setup() int {
    println("calling setup()")
    return 1
}

func main() { // happens last
    println("main")
}
