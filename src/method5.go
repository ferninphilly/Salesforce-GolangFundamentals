package main

import "fmt"

// START OMIT
type Mountain struct {
    name      string
    elevation int
}

func (m Mountain) HowBig() string {
    if m.elevation > 8800 {
        return "HUGE"
    }
    return "NORMAL"
}

func main() {
	everest := Mountain{elevation: 8848, name: "Mt. Everest"}
	fmt.Println(everest.HowBig())
}
//END OMIT
