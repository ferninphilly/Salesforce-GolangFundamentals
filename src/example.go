package example

var private = "This is not exported."
var Public = "This variable IS exported." // HL

func nothing(x int) int {
        return x - 1
}

func Something(x int) int { // HL
        return x + 1
}
