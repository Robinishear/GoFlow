package main

import "fmt"

func main() {

	//   single variable declaration

	//   var name string
	//   name = "John Doe"

	// var name string = "John Doe"

	// name := "John Doe"

	// var name = "John Doe"

	// fmt.Println("Hello, World!")
	// fmt.Println(name)

	// group variable declaration
	var (
		name     string = "John Doe"
		age      int    = 30
		email    string = "john.doe@example.com"
		lastName string = "Doe"
	)
	fmt.Println(name, age, email, lastName)

	//  multiple variable declaration
	var x, y int = 10, 20
	// x = 10
	// y = 20

	fmt.Println(x, y)
}
