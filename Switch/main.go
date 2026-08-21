package main

import "fmt"

func main() {

	day := "sunday"

	// if day == "sunday" {
	// 	fmt.Println("Today is sunday")
	// } else {
	// 	fmt.Println("Today is Not sunday")

	// }

	// tagged switch statement
	switch day {
	case "sunday":
		fmt.Println("Today is sunday")
	case "friday":
		fmt.Println("Today is friday ")
	case "weekend":
		fmt.Println("Today is weekend")

	default:
		fmt.Println("Today is Not sunday")
	}

}
