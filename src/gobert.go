package main

import "fmt"

type ByteSize float64
  
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
    fmt.Printf("%d %d %d %d\n", int(KB), int(MB), int(GB), int(TB))
}
