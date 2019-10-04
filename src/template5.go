package main

import "os"
import "text/template"

func main() {
    t := template.New("")
    // can use variables to store results and reuse them
    t.Parse(`{{ $First := .FirstName | printf "%q"}}
             {{ $Last  := .LastName  | printf "%q"}}
             Normal:  {{$First}} {{$Last}}
             Reverse: {{$Last}} {{$First}}`)
    t.Execute(os.Stdout, struct {
        FirstName string
        LastName  string
    }{
        "Dave",
        "Wade-Stein",
    })
}
