package main

import "fmt"

func main() {
	m := make(map[string]int)
	// Set a map
	m["k1"] = 4
	m["k2"] = 7
	fmt.Println(m)

	// Get a value for a key
	fmt.Println(m["k2"])

	// len of map
	fmt.Println("Len of map is : ", len(m))

	// check if a key is present in a map
	_, prs := m["k3"]
	fmt.Println(" prs ", prs)

	//declare and init a new map
	n := map[string]int{"Lundi": 1, "Mardi": 2, "Samedi": 6}
	fmt.Println("dcl is : ", n["Samedi"])
}
