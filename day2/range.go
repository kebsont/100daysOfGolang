package main

import "fmt"

func main() {

	sl := []string{"Lundi", "Mardi", "Mercredi", "Jeudi"}

	for i, v := range sl {
		fmt.Print(v)
		fmt.Printf(" is %d ieme day\n", i+1)
	}

	my_map := map[string]string{"Name": "Salimata", "age": "43", "city": "Dakar"}
	for k, v := range my_map {
		fmt.Printf("%s => %s\n", k, v)
	}

	for k := range my_map {
		fmt.Println(" Key: ", k)
	}

}
