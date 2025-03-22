package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

// declare an interface for the Client struct
// this interface is called Person and it has two methods: Activate and Desactivate
// the methods are defined in the Client struct
type Person interface {
	Activate()
	Desactivate()
}

// implement the Activate method for the Client struct
func (c Client) Activate() {
	fmt.Println("Client activated")
}

// implement the Desactivate method for the Client struct
func (c Client) Desactivate() {
	fmt.Println("Client desactivated")
}

// create another type struct
type Admin struct {
	Name string
}

// implement the Activate method for the Admin struct
func (a Admin) Activate() {
	fmt.Println("Admin activated")
}

// implement the Desactivate method for the Admin struct
func (a Admin) Desactivate() {
	fmt.Println("Admin desactivated")
}

func Shutdown(p Person) {
	p.Desactivate()
}

func main() {
	// declare a variable of type Person
	var person Person
	var admin Person

	// assign a value to the variable
	person = Client{
		Name:   "John Doe",
		Age:    30,
		Active: true,
	}

	admin = Admin{
		Name: "Admin",
	}

	// call the Shutdown function with the person and admin variables
	// because both are of type Person, the Shutdown function can be called with both
	// this is because the Shutdown function expects a parameter of type Person
	// and both person and admin implement the Person interface
	Shutdown(person)
	Shutdown(admin)
}
