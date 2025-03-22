package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	// declare a function with a closure
	closureSumMultiplyByTwo := func(a, b int) int {
		return sum(a, b) * 2
	}

	fmt.Println(closureSumMultiplyByTwo(1, 2))
}
