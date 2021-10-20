package main

import "fmt"

func main() {
	var a, b string = "initial ", "declaration"
	fmt.Println(a, b)
	var c int
	fmt.Println("c is declared but not initialised: ", c)
	f := "apple"
	fmt.Println("declared & init : ", f)
}
