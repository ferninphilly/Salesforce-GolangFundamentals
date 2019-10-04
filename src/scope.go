package main

import "fmt"

func main() {
    f := 1
    // i defined only inside the for block
    for i := 1; i < 2; i++ {
        fmt.Println("i is in scope here", i)
    }
    // fmt.Println(i) => i NOT in scope here
    if j := 4.5; j > 10 {
        fmt.Println("j is in scope here:", j)
    } else {
        fmt.Println("j also is in scope here:", j)
    }
    // fmt.Println(j) => j NOT in scope here
    if f == 1 {
        f := "new f, inside this block, hides the outer f"
        fmt.Println("in the block, f is", f)
    }
    fmt.Println(f)
}
