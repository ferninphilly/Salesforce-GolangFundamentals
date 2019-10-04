package main

// START OMIT
import "text/template"
import "os"

type Joke struct {
    Who, Punchline string
}
func main() {
    t := template.New("Knock Knock Joke") // create a template
    text := `Knock Knock
                Who's there?
{{.Who}}
                {{.Who}} who?
{{.Punchline}}
`
    t.Parse(text) // parse the text to generate an "executable" template
    jokes := []Joke{
        {"Cash", "No thanks, I'll have peanuts!"},
        {"Cow says", "No, cow says moo!"},
    }
    for _, joke := range jokes {
        t.Execute(os.Stdout, joke) // execute the template
        println()
    }
}
// END OMIT
