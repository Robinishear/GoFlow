package main

import "fmt"

// type user struct {
// 	name      string
// 	age       int
// 	isLogging bool
// 	great     func()
// }

type user struct {
	name      string
	age       int
	isLogging bool
}

func main() {
	// users := user{
	// 	name:      "Robin",
	// 	age:       25,
	// 	isLogging: false,
	// }
	// users.great = func() {
	// 	fmt.Println(users.name, users.age, users.isLogging)
	// }
	// // fmt.Println(users)
	// users.great()

	users := user{
		name:      "Robin",
		age:       25,
		isLogging: false,
	}

	users.great()
	users.login()

	// pointerUsers := &users
	// pointerUsers.login()

	fmt.Printf("%+v", users)

}

func (u *user) login() {
	fmt.Println("Logging now")
	u.isLogging = true
}

func (u user) great() {
	fmt.Println("Hello! My name is ", u.name, "My age is ", u.age, "and I logging ", u.isLogging)
}
