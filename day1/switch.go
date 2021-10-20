package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	//basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Use comma in switch and the default case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch like a if/else
	switch {
	case time.Now().Hour() < 12:
		fmt.Print("Bonjour ")
	default:
		fmt.Print("Bonsoir ")
	}
	fmt.Println("cher dev")

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case string:
			fmt.Println("I'm a string")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("I don't know this type. It seems it's a %T", t)
		}
	}
	whatAmI(1)
	whatAmI("hello")
	whatAmI(true)
	whatAmI(5.5)

}
