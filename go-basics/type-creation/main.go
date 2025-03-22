package main

type ID int

func main() {
	var a ID = 1
	println(a)

	var b ID = 2
	println(b)

	b = a // error: cannot use a (type ID) as type ID in assignment
	println(b)
}
