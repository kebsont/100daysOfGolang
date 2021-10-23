package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {

		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(5, 6)
	sum(4, 3)
	// fmt.Println("Sum : ", sum(5, 5, 6))
}
