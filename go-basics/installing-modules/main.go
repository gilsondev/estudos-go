package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Using uuid module for create a new UUID value")

	fmt.Println("Generating a new UUID value")
	uuid := uuid.New()
	fmt.Println("UUID Value:", uuid)
}
