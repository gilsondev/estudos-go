package main

import "fmt"

// Define a struct type called Person
type Person struct {
	Name string
	Age  int
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}

// NewPerson is a constructor function that returns a pointer to a new Person struct
// This is a common pattern in Go to create new instances of a struct
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// Create a new Person struct
	person := NewPerson("John", 30)

	// Print the person struct
	fmt.Println("Person:", person)

	// Print the person struct using a pointer
	fmt.Println("Person:", &person)

	// Change the name of the person
	fmt.Println("Changing name of person to Jane")
	person.ChangeName("Jane")

	// Print the person struct again
	fmt.Println("Person:", person)

}
