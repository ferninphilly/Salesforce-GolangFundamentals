package main

import "os"
import "fmt"
// START OMIT
import "text/template"
type Person struct {
    Name string
}
type Family struct {
    Father, Mother Person
}
func main() {
    family := Family{
        Father:      Person{"Tarzan"},
        Mother:      Person{"Jane"},
    }
    t := template.New("Father")
    text := "The father's name is {{.Father.Name}}"
    if _, err := t.Parse(text); err != nil {
        fmt.Printf("Failed to parse %s (%s)\n", text, err)
    } else {
        t.Execute(os.Stdout, family)
    }
}
// END OMIT
