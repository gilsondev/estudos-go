package main

import "fmt"

func main() {
	// For in Go is the only loop construct
	// It can be used in three different ways:
	// 1. for init; condition; post {
	// 2. for condition {
	// 3. for {
	// }

	// 1. for init; condition; post {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// 2. for condition {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// 3. for {
	i := 0
	for {
		fmt.Println("i:", i)
		i++
		if i >= 10 {
			break
		}
	}

	// Nested loops
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}

	// Using labels to break outer loops
outerLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 && j == 5 {
				break outerLoop
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// Using labels to continue outer loops
anotherLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 && j == 5 {
				continue anotherLoop
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// Range is used to iterate over a slice, array, map, string, etc.
	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		fmt.Println("index:", index, "number:", number)
	}

	// Range is also used to iterate over a map
	person := map[string]string{
		"name": "John",
		"age":  "20",
	}
	for key, value := range person {
		fmt.Println("key:", key, "value:", value)
	}

	// Range is also used to iterate over a string
	message := "Hello, World!"
	for index, char := range message {
		fmt.Println("index:", index, "char:", string(char))
	}
}
