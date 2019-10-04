package main

import "fmt"

type Mountain struct {
    name      string
    elevation int
}

// START OMIT
type Climb struct {
    Mountain
    climber  string
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
	arjun := Climb{Mountain{"K2", 8111}, "Arjun Climber"}
	fmt.Println(arjun.HowBig())
}
// END OMIT
