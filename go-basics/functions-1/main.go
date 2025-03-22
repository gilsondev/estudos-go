package main

import (
	"errors"
	"fmt"
)

// declare a function
func sayHello() {
	fmt.Println("Hello, World!")
}

// declare a function with a parameter
// PS.: You can omit the type of the parameters if they are of the same type
// Example: `func sum(a, b int) int`
func sum(a int, b int) int {
	return a + b
}

// declare a function with a multiple return values
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// In Go, we dont have exception handling, but we have error handling
// We use the `error` type to handle errors
func sumWithError(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be positive")
	}
	return a + b, nil
}

func main() {
	// call the function
	sayHello()

	// call the function with a parameter
	fmt.Println(sum(1, 2))

	// call the function with a multiple return values
	fmt.Println(divide(10, 3))

	// create a variable that will be a inner function
	innerSum := func(a int, b int) int {
		return a + b
	}

	// call the inner function
	fmt.Println(innerSum(1, 2))

	// call the function with a error
	result, err := sumWithError(10, -2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(result)
}
