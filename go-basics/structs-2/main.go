package main

import "fmt"

// declare a struct
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	// embed a struct into another struct
	// this is equivalent to:
	// Address Address
	Address
}

func main() {
	// declare a variable of type Client and initialize it with a value
	client := Client{
		Name:   "John Doe",
		Age:    30,
		Active: true,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
			Zip:    "12345",
		},
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", client.Name, client.Age, client.Active)

	// change the value of the Active field
	fmt.Println("Changing the value of the Active field")
	client.Active = false
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", client.Name, client.Age, client.Active)

	// access the fields of the embedded struct
	fmt.Printf("Street: %s, City: %s, State: %s, Zip: %s\n", client.Street, client.City, client.State, client.Zip)
}
