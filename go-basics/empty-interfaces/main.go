package main

import "fmt"

// What is an empty interface?
// An empty interface is an interface that has no methods
// It is represented as interface{}
// It is used to represent a set of methods, a type that has no methods
// This approach is used to create a type that can be used to represent any type
// like a generic type. But in recent versions of Go, this approach is discouraged
// and the use of generics is recommended instead.

// Define an empty interface
type EmptyInterface interface{}

func main() {
	// Declare a variable of type EmptyInterface of type string
	// and int are both valid types for the empty interface
	var emptyString EmptyInterface = "Hello, World!"
	var emptyInt EmptyInterface = 10

	// Print the values of the empty interface
	fmt.Println(emptyString)
	fmt.Println(emptyInt)

	// Print the type of the empty interface
	fmt.Printf("Type of emptyString: %T\n", emptyString)
	fmt.Printf("Type of emptyInt: %T\n", emptyInt)
}
