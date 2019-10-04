// Given the following type:
//
//	type IPAddr [4]byte
//
// 1. Define a map which maps hostnames (string) to `IPAddr`
// 2. Add a few hostnames to your map, e.g., "localhost.com"
// (the addresses don't have to be correct)
// 3. Add a `Stringer` method to `IPAddr` so that it is printed
// in dotted notation (e.g., "127.0.0.1")

package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hostnames := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"google.com": {1, 2, 3, 4},
	}

	for hostname, addr := range hostnames {
		fmt.Println(hostname, addr)
	}
}
