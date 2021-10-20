package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 32
	fmt.Println("set ", a)
	fmt.Println("get a[4]", a[4])

	new_array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl : ", new_array)
	var twoD [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
