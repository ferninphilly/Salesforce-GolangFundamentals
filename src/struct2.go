package main

import "fmt"

type mountain struct {
    name      string
    elevation int
}

// START OMIT
func main() {
    everest := mountain{elevation: 8848, name: "Mt. Everest"}
    fmt.Printf("%#v\n", everest)
    fmt.Println(mountain{name: "Flat"})
    fmt.Println(&mountain{name: "Long's Peak", elevation: 4346})
    nd := mountain{name: "Nanda Devi", elevation: 7815}
    fmt.Println(nd.name)
    sp := &nd
    fmt.Println(sp.elevation)
    sp.elevation++
    fmt.Println(sp.elevation)
}
// END OMIT
