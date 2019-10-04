package main

import "fmt"
import "encoding/json"

// START OMIT
type response1 struct {
    Page   int
    Fruits []string
}

func main() {
    res1D := &response1{ // Only exported fields are included in the encoded...
        Page:   1,       // ...output and will by default use those names as keys
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D) // HL
    fmt.Println(string(res1B)) 
    // Decode JSON into custom data types. Adds additional type-safety and
    // eliminates need for type assertions when accessing decoded data.
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response1{}
    json.Unmarshal([]byte(str), &res) // HL
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
}
// END OMIT
