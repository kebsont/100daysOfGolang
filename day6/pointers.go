package main

import "fmt"

func main() {
	x := 5
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)
	natal := "France"
	changeValue(&natal)
	fmt.Println(natal)
	living := "Belgique"
	CopyChangingValue(living)
	fmt.Println(living)
}

func changeValue(country *string) {
	*country = "Senegal"
}
func CopyChangingValue(country string) {
	country = "Senegal"
}
