package main

// START OMIT
func myfunc(f float64) int {
    return int(f)
}

func main() {
    x := 3
    y := myfunc(x)
    print(y)
}
// END OMIT
