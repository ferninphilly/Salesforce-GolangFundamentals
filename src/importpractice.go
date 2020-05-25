package main

import (
	"fmt"
	"hello"
)

func main() {
	fmt.Println(hello.ExportConstant)
	fmt.Println(hello.DoubleIt(5, 6))
}
