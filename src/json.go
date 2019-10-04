package main

// START OMIT
import "fmt"
import "encoding/json"

func main() {
    b, i, f, s := true, 1, 2.34, "gopher" // basic types
    bolB, _ := json.Marshal(b) // turn basic types into bytes // HL
    intB, _ := json.Marshal(i) // HL
    fltB, _ := json.Marshal(f) // HL
    strB, _ := json.Marshal(s) // HL
    fmt.Printf("%#v %s\n", bolB, string(bolB))
    fmt.Printf("%#v %s\n", intB, string(intB))
    fmt.Printf("%#v %s\n", fltB, string(fltB))
    fmt.Printf("%#v %s\n", strB, string(strB))
    _ = json.Unmarshal(bolB, &b) // ignore errors for now // HL
    _ = json.Unmarshal(intB, &i) // HL
    _ = json.Unmarshal(fltB, &f) // HL
    _ = json.Unmarshal(intB, &s) // HL
    fmt.Printf("%#v %#v %#v %#v\n", b, i, f, s)
}
// END OMIT
