package main

import "fmt"

func makeCoffee(kind string) {
	fmt.Printf("Coffee %s is ready.... \n", kind)
}

func main() {
	makeCoffee("black coffee")
	makeCoffee("white coffee")
	makeCoffee("Latte")
	makeCoffee("thank you for coffee")
}

// func main() {

// 	//   single variable declaration

// 	//   var name string
// 	//   name = "John Doe"

// 	// var name string = "John Doe"

// 	// name := "John Doe"

// 	// var name = "John Doe"

// 	// fmt.Println("Hello, World!")
// 	// fmt.Println(name)

// 	// group variable declaration
// 	// var (
// 	// 	name     string = "John Doe"
// 	// 	age      int    = 30
// 	// 	email    string = "john.doe@example.com"
// 	// 	lastName string = "Doe"
// 	// )
// 	// fmt.Println(name, age, email, lastName)

// 	//  multiple variable declaration
// 	// var x, y int = 10, 20
// 	// x = 10
// 	// y = 20

// 	// fmt.Println(x, y)

// 	// var age int
// 	// fmt.Println(age) // default value is 0

// 	// var name string
// 	// fmt.Println(name) // default value is ""

// 	// var isAdmin bool
// 	// fmt.Println(isAdmin) // default value is false

// 	// var amount float64
// 	// fmt.Println(amount) // default value is 0.0

// }
