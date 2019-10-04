// Write a program which takes two command line args representing integers
// and prints out their sum
package main

import (
        "os"
        "fmt"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])	
	fmt.Println("sum of", os.Args[1], "and", os.Args[2], "is", num1+num2)
}
