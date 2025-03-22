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

func (c Client) Desactive() {
	c.Active = false
	fmt.Println("Client desactivated")
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

	fmt.Println("Desactivating client")
	client.Desactive()
}
