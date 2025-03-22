package main

import "fmt"

// Type assertion is a way to check the type of an interface
// and to get the value of the interface as the specific type

// Create a new empty interface
type EmptyInterface interface{}

func main() {
	// Declare a variable of type EmptyInterface
	var myVar EmptyInterface = "Hello, World!"

	// Print the value of the variable
	fmt.Println("Creating a new variable of type EmptyInterface")
	fmt.Println(myVar)

	// Print the type of the variable
	fmt.Printf("Type of myVar: %T\n", myVar)

	// Type assertion to get the value of the interface as the specific type
	fmt.Println("Type assertion to get the value of the interface as the specific type: string")
	myString, ok := myVar.(string)

	// Print the value of the variable
	fmt.Printf("myString: %v, ok: %v\n", myString, ok)

	// Print the type of the variable
	fmt.Printf("Type of myString: %T\n", myString)

	// Type assertion to get the value of the interface as the specific type
	fmt.Println("Type assertion to get the value of the interface as the specific type: int")
	myInt, ok := myVar.(int)

	// Print the value of the variable
	fmt.Printf("myInt: %v, ok: %v\n", myInt, ok)

	// If you try to type assert to a type that is not the same as the type of the value,
	// and not read the second return value to handle the error, you will get a panic
	fmt.Println("If you try to type assert to a type that is not the same as the type of the value, you will get a panic")
	myBool := myVar.(bool)

	// Print the value of the variable
	fmt.Printf("myBool: %v\n", myBool)

}
