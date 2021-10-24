package main

import "fmt"

func idontknow() func(x int) int {
	count := 0
	return func(x int) int {
		for i := 0; i < 5; i++ {
			count++
		}
		return x + count
	}

}

func main() {
	res := idontknow()

	fmt.Println(res(5))
	fmt.Println(res(7))
	fmt.Println(res(9))
}
