package main

import (
	"fmt"
	"math"
)

func main() {
	// if with condition
	a := 4
	b := int(math.Sin(float64(a)))
	if num := 9; num < b {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}

// There is no ternary if in Go.
