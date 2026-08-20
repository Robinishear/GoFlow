package main

import "fmt"

// import "fmt"

func main() {

	// Ananymous Function

	// coffeeOrder := func() {
	// 	fmt.Println("Coffee order placed")
	// }
	// coffeeOrder()

	func() {
		fmt.Println("Coffee order placed")

	}()
}
