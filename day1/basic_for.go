package main

import (
	"fmt"
)

func main() {
	// var chaine string = "hello"
	arr := [7]int{2, 5, 3, 77, 23, 1, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Println()

	for {
		fmt.Println("Hey you")
		fmt.Println("yes")
		break
		fmt.Println("Normally I won't appear")
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
