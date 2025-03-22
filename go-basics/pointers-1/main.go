package main

import "fmt"

func main() {
	// Memory -> Address -> Value
	a := 10

	// Print the value of a
	fmt.Println("Value of a:", a)

	// Print the memory address represented by the & operator
	fmt.Println("Address of a:", &a)

	// Declare a pointer to a
	var b *int

	// Assign the memory address of a to b
	b = &a

	// Print the value of b (dereferencing the pointer)
	// *b is the value at the memory address stored in b
	fmt.Println("Value of b:", *b)

	// Print the memory address of b
	fmt.Println("Address of b:", &b)

	// Change the value of a using the pointer
	fmt.Println("Value of a before change:", a)
	*b = 20

	// Print the value of a again
	fmt.Println("Value of a after change:", a)

	// Print the value of b again
	fmt.Println("Value of b:", *b)

}
