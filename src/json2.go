package main

// START OMIT
import "fmt"
import "encoding/json"

func main() {
    // Slices/maps encode to JSON arrays and objects as you'd expect
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD) // HL
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD) // HL
    fmt.Println(string(mapB))
    // Decoding JSON into custom data types. This has the advantages
    // of adding additional type-safety to our programs and eliminating
    // the need for type assertions when accessing the decoded data.
    var slice = []string{}
    json.Unmarshal([]byte(slcB), &slice) // HL
    fmt.Printf("%#v\n", slice)
    var unmap map[string]int
    json.Unmarshal([]byte(mapB), &unmap) // HL
    fmt.Printf("%#v\n", unmap)
}
// END OMIT
