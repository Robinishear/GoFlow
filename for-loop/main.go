package main

import "fmt"

func makeCoffee(x int) {
	fmt.Println("making coffee............!!!", x)
}

func main() {

	// for i := 0; i <= 10; i++ {   // for initialization; condition; increment/decrement
	// 	// fmt.Println(i)
	// 	makeCoffee(i)
	// }

	// while styled loop

	i := 1

	for i <= 5 {
		makeCoffee(i)
	}
}

// i := 1, true, run the body, increment
// i = 2, true, run the body, increment
// i = 3, true, run the body, increment
// i = 4, true, run the body, increment
// i = 5, true, run the body, increment
// i = 6, false, stop
