package main

import (
	"fmt"
	"math"
)

func somme(a int, b int) int {
	return a + b
}

func racine(a float64) float64 {
	return float64(math.Sqrt(a))
}

// multiple returns
func multiple_returns(a int, b float64) (int, int) {
	return somme(a, int(b)), int(racine(b))
}

func main() {
	a := 5
	b := 7
	c := 9.0
	fmt.Println("Somme = ", somme(a, b))
	fmt.Println("Racine carre = ", racine(c))
	i, j := multiple_returns(a, c)
	fmt.Println("Multiple Returns = ", i, j)

}
