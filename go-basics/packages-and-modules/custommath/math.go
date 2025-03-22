package custommath

// Sum is a function that use modifier public to be used outside the package
func Sum[T int | float64](a, b T) T {
	return anotherSum(a, b)
}

// But this sum function is not exported, so it is not available outside the package
// and this is possible because the function name starts with a lowercase letter
func anotherSum[T int | float64](a, b T) T {
	return a + b
}

// This rule is valid for variables, constants, types, etc.
// If the name starts with a lowercase letter, it is not exported
// If the name starts with an uppercase letter, it is exported
// This is called "name exportation"

// In structs, the same rule applies to the fields
// If the field starts with a lowercase letter, it is not exported
// If the field starts with an uppercase letter, it is exported

// In interfaces, the same rule applies to the methods
// If the method starts with a lowercase letter, it is not exported
