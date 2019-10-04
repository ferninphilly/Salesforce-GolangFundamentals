package main

import "fmt"

// START OMIT
func main() {
    // To create an empty map, use make
    m := make(map[string]int)
    // Set key/value pairs
    m["key1"] = 14
    m["key2"] = -2
    m["key3"] = 0
    fmt.Println(m)
    fmt.Println("len:", len(m)) // HL
    delete(m, "key2") // HL
    fmt.Println(m)

    // Optional second return value indicates if key was in map.
    // Used to disambiguate between missing keys and zero values.
    _, present := m["key2"]
    fmt.Println("key2 present?", present)
}
// END OMIT
