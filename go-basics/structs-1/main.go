package main

import "fmt"

// declare a struct
type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// declare a variable of type Client
	client := Client{
		Name:   "John Doe",
		Age:    30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", client.Name, client.Age, client.Active)

	// change the value of the Active field
	fmt.Println("Changing the value of the Active field")
	client.Active = false
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", client.Name, client.Age, client.Active)
}
