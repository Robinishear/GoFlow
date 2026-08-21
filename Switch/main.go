package main

import "fmt"

func main() {

	day := "weekend"

	// if day == "sunday" {
	// 	fmt.Println("Today is sunday")
	// } else {
	// 	fmt.Println("Today is Not sunday")

	// }

	switch day {
	case "sunday":
		fmt.Println("Today is sunday")
	case "friday":
		fmt.Println("Today is friday ")
	case "weekend":
		fmt.Println("Today is weekend")

	}

}
