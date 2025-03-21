package main

import (
	"fmt"

	"github.com/gilsondev/estudos-go/go-expert-packaging/example-4/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
