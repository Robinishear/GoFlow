package main

import "fmt"

func main() {
	// var orders = [6]int{10, 20, 30, 40, 50, 60}

	// fmt.Println(orders[:3])

	// slice := orders[2:5] // start from index 3 and end at index 4(exclusive of 4)
	// fmt.Println(slice)

	// slice[0] = 500
	// fmt.Println(slice)
	// fmt.Println(orders)

	// fmt.Println("len:", len(slice))
	// fmt.Println("cap:", cap(slice))

	// slice literal
	var slice = []int{
		100,
		200,
		300,
		400,
		500,
		600,
		700,
		800,
		900,
		1000,
	}
	slice[0] = 50
	slice[1] = 150
	slice[2] = 250
	slice[3] = 350
	slice[4] = 450
	slice[5] = 550
	slice[6] = 650
	slice[7] = 750
	slice[8] = 850
	slice[9] = 950

	slice = append(slice, 22)

	fmt.Println("The slice:", slice)
	fmt.Println("Length of the slice:", len(slice))
	fmt.Println("Capacity of the slice:", cap(slice))

}
