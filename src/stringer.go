package main
 
import "fmt"

// START OMIT
type Person struct {
    Name, AlmaMater string
}

func (p Person) String() string {
    return fmt.Sprintf("%v (Alma Mater: %v)", p.Name, p.AlmaMater)
}

func main() {
    rs := Person{"Ranveer Singh", "Indiana University"}
    ab := Person{"Amitabh Bachchan", "Sherwood College"}
    fmt.Printf("%s\n%s\n", rs, ab)
}
// END OMIT
