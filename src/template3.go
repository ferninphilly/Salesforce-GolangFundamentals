package main

import "math"
import "os"
import "text/template"

func main() {
    t := template.New("")
    t.Parse(`15 digits of π: {{printf "%.15f\n" .}}`)
    t.Execute(os.Stdout, math.Pi)
    // Must() is a helper method that panics if an error occurs
    t2 := template.Must(template.New("").Parse(`15 digits of π: {{printf "%.15f\n" .}}`))
    t2.Execute(os.Stdout, math.Pi)
}
