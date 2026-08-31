package main

import "fmt"

func main() {
	// MyMap := make(map[string]int)

	// MyMap["robin"] = 15
	// MyMap["anirban"] = 18
	// MyMap["adit"] = 19

	// fmt.Println(MyMap)

	//
	MyMap := map[string]string{
		"name":    "robin",
		"success": "ok",
	}

	delete(MyMap, "success")
	fmt.Println(MyMap)
}
