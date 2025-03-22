package main

const a = "Hello, World!"

// Global scope
var b bool // default value is false
var (
	c int         // default value is 0
	d string      // default value is ""
	e int    = 10 // declaration and assignment
)

func main() {
	println(a)

	b = true // assignment
	println(b)

	c = 10 // assignment
	println(c)

	d = "Hello, World!" // assignment
	println(d)

	println(e)

	// declaration and assignment and local scope. A shorthand for var f int = 10
	// PS 1: You can't assign a different type to the variable
	f := 10
	println(f)

	var g string // declaration with default value
	println(g)
}
