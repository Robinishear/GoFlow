package main

import "fmt"

// import "fmt"

func main() {

	// Ananymous Function

	coffeeOrder := func() {
		fmt.Println("Coffee order placed")
	}
	coffeeOrder()

	// Immediately Invoked Function Expression (IIFE)
	func(CoffeeType string) {
		fmt.Printf("Coffee order placed  %s.........", CoffeeType)

	}("Latte")
}
