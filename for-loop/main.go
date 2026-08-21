package main

import "fmt"

func makeCoffee(x int) {
	fmt.Println("making coffee............!!!", x)
}

func main() {

	for i := 0; i <= 10; i++ {
		// fmt.Println(i)
		makeCoffee(i)
	}

}
