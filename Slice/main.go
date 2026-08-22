package main

import "fmt"

func main() {
	var orders = [6]int{10, 20, 30, 40, 50}

	// fmt.Println(orders[:3])

	Slice := orders[3:4]
	fmt.Println(Slice)

}
