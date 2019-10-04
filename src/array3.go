package main

import "fmt"

// START OMIT
func main() {
    var twoD [3][5]int // 2-dimensional array
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            twoD[i][j] = (i + 1) * (j + 1)
        }
    	fmt.Println("Row", i, twoD[i])
    }
    fmt.Println("\nAll at once:", twoD)
}
// END OMIT
