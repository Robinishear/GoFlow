package main

import "fmt"

// array type is [size]type and array is value type array length and capacity is same
func main() {

	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// fmt.Println(numbers)

	// get array length
	// fmt.Println("Length", len(numbers))

	// fmt.Println("Index is: ", numbers[4], "Type is: ", numbers[3])\

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index is: ", i, "Value is: ", numbers[i])
	}

}
