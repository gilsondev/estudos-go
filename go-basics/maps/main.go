package main

import "fmt"

func main() {
	// declare a map of string to int
	salaries := map[string]int{
		"John": 10000,
		"Jane": 20000,
		"Jim":  30000,
	}

	// print the map
	fmt.Println(salaries)

	// removing an item from the map
	delete(salaries, "Jim")
	fmt.Println(salaries)

	// adding an item to the map
	salaries["John"] = 11111
	fmt.Println(salaries)

	// create a map with make
	fmt.Println("Creating a map with make")
	salaries2 := make(map[string]int)
	fmt.Println(salaries2)

	// add an item to the map
	salaries2["John"] = 11111
	fmt.Println(salaries2)

	// create a map with shorthand notation
	fmt.Println("Creating a map with shorthand notation")
	salaries3 := map[string]int{
		"John": 11111,
		"Jane": 22222,
		"Jim":  33333,
	}
	fmt.Println(salaries3)

	// iterate over a map
	fmt.Println("Iterating over a map")
	for key, value := range salaries3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Iterate over a map with range and blank identifier
	fmt.Println("Iterating over a map with range and blank identifier")
	for _, value := range salaries3 {
		fmt.Println(value)
	}

	// check if a key exists in a map
	fmt.Println("Checking if a key exists in a map")
	value, exists := salaries3["John"]
	fmt.Println(value, exists)

}
