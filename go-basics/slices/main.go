package main

import "fmt"

func main() {
	// declare a slice of int
	var mySlice []int

	// assign values to the slice
	fmt.Println("Assigning values to the slice")
	mySlice = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// print the slice
	fmt.Println(mySlice)

	// print the size of the slice
	fmt.Println("The size of the slice is", len(mySlice))

	// print the capacity of the slice
	fmt.Println("The capacity of the slice is", cap(mySlice))

	// print the capacity of the slice using the built-in function
	fmt.Println("The capacity of the slice using the built-in function 'len' is", len(mySlice))

	fmt.Println("Slicing the slice")
	fmt.Println("mySlice cap:", cap(mySlice))
	fmt.Printf("Using mySlice[2:]: len=%d cap=%d %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])
	fmt.Printf("Using mySlice[:2]: len=%d cap=%d %v\n", len(mySlice[:2]), cap(mySlice[:2]), mySlice[:2])

	// appending to a slice
	fmt.Println("Appending to a slice")
	mySlice = append(mySlice, 110)
	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)

	// appending to a slice with a slice
	fmt.Println("Appending to a slice with a slice")
	mySlice = append(mySlice, mySlice...)
	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)

}
