package main

import "fmt"

// type PremiumUser struct {
// 	phone   int
// 	address string
// 	plan    string
// 	website string
// }

// // struct is a collection of fields
// type user struct {
// 	name     string
// 	email    string
// 	password string
// 	metaInfo PremiumUser
// }

type user struct {
	name string
	age  int
	role string
}

// Anonymous struct
func main() {
	// jon := user{name: "jon",
	// 	email:    "jon@gmai.com",
	// 	password: "JIokj#@$55%^^&",

	// 	metaInfo: PremiumUser{phone: 1234567890,
	// 		address: "123 Main St",
	// 		plan:    "Premium",
	// 		website: "example.com"}}
	// fmt.Printf("%+v", jon)

	newUser := func(name string, age int, role string) user {
		return user{
			name: name,
			age:  age,
			role: role,
		}
	}

	users := newUser("jon", 25, "admin")
	fmt.Printf("%+v", users)
}
