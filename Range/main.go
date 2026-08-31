package main

import "fmt"

func main() {
	// MyMap := map[string]string{
	// 	"Name":      "robin",
	// 	"success":   "ok",
	// 	"language":  "go",
	// 	"rating":    "5",
	// 	"hours":     "25",
	// 	"price":     "100",
	// 	"country":   "Bangladesh",
	// 	"state":     "Dhaka",
	// 	"city":      "Dhaka",
	// 	"zip":       "1200",
	// 	"street":    "Dhaka",
	// 	"house":     "Dhaka",
	// 	"apartment": "Dhaka",
	// 	"floor":     "1",
	// 	"wing":      "A",
	// 	"flat":      "101",
	// 	"building":  "Dhaka",
	// 	"road":      "Dhaka",
	// 	"area":      "Dhaka",
	// 	"village":   "Dhaka",
	// 	"district":  "Dhaka",
	// 	"division":  "Dhaka",
	// }
	// for key, value := range MyMap {
	// 	fmt.Printf("%v : %v \n", key, value)
	// }

	name := "robinryan"

	for i, value := range name {
		fmt.Println(i, value)
	}

}
