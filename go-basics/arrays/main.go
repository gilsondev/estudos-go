package main

import "fmt"

func main() {
	// declare an array of 3 positions
	var myArray [3]int

	// assign values to the array
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	// print the array
	fmt.Println(myArray)

	// print size of the array
	fmt.Println(len(myArray))

	// print the type of the array
	fmt.Printf("The type of myArray is %T", myArray)

	// access the value of the array
	fmt.Println(myArray[0])

	// change the value of the array
	myArray[0] = 10
	fmt.Println(myArray)

	// Using range to iterate over the array
	for index, value := range myArray {
		fmt.Printf("The value of the array at index %d is %d\n", index, value)
	}
}
