package main

import "fmt"

func main() {
	//  age 18 theke beshi hoile true hobe na hoi false

	// age := 20
	// isAdult := age >= 18
	// fmt.Print(isAdult)
	// if age >= 18 {
	// 	fmt.Println("You are eligible for voting")
	// }

	//  grading system
	score := 100

	if score >= 80 {
		fmt.Println("Grade A", score)
	} else if score >= 70 {
		fmt.Println("Grade B", score)
	} else if score >= 60 {
		fmt.Println("Grade C", score)
	} else if score >= 50 {
		fmt.Println("Grade D", score)
	} else {
		fmt.Println("Grade F sala fa fa", score)
	}
}
