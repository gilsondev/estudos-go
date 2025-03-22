package main

import "fmt"

func main() {
	// Conditionals in Go are similar to other languages, but there are some differences
	// For example, the if statement does not need parentheses
	if true {
		fmt.Println("The condition is true")
	}

	// The else statement is used to execute a block of code if the condition is false
	if false {
		fmt.Println("The condition is false")
	} else {
		fmt.Println("The condition is false")
	}

	// The else if statement is used to execute a block of code if the condition is false
	if false {
		fmt.Println("The condition is false")
	} else if true {
		fmt.Println("The condition is true")
	} else {
		fmt.Println("The condition is false")
	}

	// The switch statement is used to execute a block of code based on a condition
	switch true {
	case true:
		fmt.Println("The condition is true")
	case false:
		fmt.Println("The condition is false")
	}

	// The switch statement can also be used with a variable
	switch "test" {
	case "test":
		fmt.Println("The condition is true")
	case "test2":
		fmt.Println("The condition is false")
	}
}
