package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("======= Sum Program =======")

	var a, b int
	fmt.Println("Enter the first number:")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&b)
	fmt.Println("The sum of the numbers is:", Sum(a, b))
}
