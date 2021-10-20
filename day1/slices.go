package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "h"
	s[1] = "e"
	s[2] = "y"
	fmt.Println("set ", s)
	fmt.Println("get ", s[2])
	//append
	s = append(s, "")
	s = append(s, "Mouss")
	fmt.Println("after append", s)
	//copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied slice is: ", c)
	//slice operation
	l := s[2:5]
	fmt.Println("Sliced : ", l)
	l = s[:3]
	fmt.Println("Sliced :  s[:3]", l)
	l = s[3:]
	fmt.Println("Sliced : s[3:]", l)

	//declare and init
	t := []string{"h", "e", "l", "l", "o"}
	fmt.Println("dclr", t)

	// 2D slices
	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println("twoD: ", twoD)
}
