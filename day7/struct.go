package main

import "fmt"

type car struct {
	brand string
	color string
	miles int
}

func newCar(brand string) *car {
	cx3 := car{brand: brand}
	cx3.color = "red"
	cx3.miles = 143000
	return &cx3

}

func main() {
	fmt.Println(car{brand: "Mitsubishi", color: "white", miles: 0})
	fmt.Println(car{brand: "Toyota"})
	fmt.Println(&car{brand: "Mercedes", color: "black"})
	fmt.Println(newCar("Jeep"))

	rav4 := car{brand: "Toyota", miles: 20000, color: "black"}
	fmt.Println(rav4.miles)

	rav4_2020 := &rav4
	fmt.Println(rav4_2020.miles)

	rav4_2020.color = "Red"
	fmt.Println(rav4_2020.color)
}
