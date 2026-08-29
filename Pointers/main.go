package main

import "fmt"

// func main() {
// 	a := 42
// 	p := &a //  p is store the address of a  0xe52d3b5e138

// 	a = 500

// 	// & → address of a variable
// 	// * → dereference (value from address)

// 	*p = 1000

// 	fmt.Println("a:", a)
// 	fmt.Println("p:", p)
// 	fmt.Println("p:", *p)
// }

func change(x *int) {
	*x = 100
}

func main() {
	num := 10

	change(&num)

	fmt.Println(num)
}
