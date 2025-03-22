package main

import "fmt"

// declare a function with a variadic parameter
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// call the function with a variadic parameter
	fmt.Println(sum(1, 2, 3, 4, 5))
}
