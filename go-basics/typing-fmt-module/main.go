package main

import "fmt"

type ID int

func main() {
	var a ID = 1
	fmt.Println(a)

	// print the type of a
	fmt.Printf("The type of a is %T", a)

}
