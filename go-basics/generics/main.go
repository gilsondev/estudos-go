package main

import "fmt"

// Generics are a way to create functions that can work with multiple types
// They are a way to create reusable, efficient, readable, and maintainable code

// Define a generic function that can work with multiple types
func Sum[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// In the Go you can use the type constraint to specify the type of the generic function
// and use some operators to specify the types. For example:
// T1 | T2 means that the function can work with T1 or T2
// T1 & T2 means that the function can work with T1 and T2
// *T means that the function can work with a pointer to a type
// []T means that the function can work with an array of a type
// map[string]T means that the function can work with a map of a type

// You can create a custom type constraint by defining a type alias
type Number interface {
	int | float64
}

// Now you can use the custom type constraint to specify the type of the generic function
func SumWithNumbers[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// It may happen that we have two custom types, and both use int for example.
// In this case, we can use the type constraint to specify the type of the generic function
type Number2 interface {
	int
}

// If we have a generic method that accepts a custom type, but also accepts int under the cloths, we need to specify for the GO, that we will open an exception with the character (~)
// For example:
// type Number interface {
// 	~int | float64
// }
// Now, the type constraint will be ~int, which means that the function can work with int or any type that is based on int

func main() {
	// Declare a map of strings to ints
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	mFloat := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}

	// Call the Sum function
	sum := Sum(m)
	sumFloat := Sum(mFloat)

	// Print the sum
	fmt.Println(sum)
	fmt.Println(sumFloat)
}
