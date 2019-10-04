package main

import (
        "fmt"
        "thing"    // singleton, i.e., only one object // HL
)

func main() {
    // try singleton "thing"
    fmt.Println(thing.ReadThing())
    thing.WriteThing(-5)
    fmt.Println(thing.ReadThing())
}

