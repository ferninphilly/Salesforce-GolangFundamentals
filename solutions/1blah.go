// Create your own package in your `$GOROOT` directory
// Be sure to export at least one function and one constant
// Build and install your package (go install)
// Write a program in different subdirectory which makes use of your package

// 1. Put this in ~/go/src/blah/blah.go
// 2. go build
// 3. go install
// 4. Now import it in a main program in another subdir
package blah

const BlahThing = 12.5

func BlahFunc() float64 {
	return BlahThing
}

func privateFunc() {
// this is private
}
