package main

import "text/template"
import "os"

func main() {
    // pipelines let you call multiple functions, as in Bash
    t := template.Must(template.New("").Parse(`{{len . | printf "%d items\n" }}`))
    t.Execute(os.Stdout, "12345")
}
