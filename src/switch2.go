package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Kolkata", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Lafayette")
	fmt.Printf("Dave works in %s, %s\n", region, continent)
}
