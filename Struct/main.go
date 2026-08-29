package main

import "fmt"

type PremiumUser struct {
	phone   int
	address string
	plan    string
	website string
}

// struct is a collection of fields
type user struct {
	name     string
	email    string
	password string
	metaInfo PremiumUser
}

func main() {
	jon := user{name: "jon",
		email:    "jon@gmai.com",
		password: "JIokj#@$55%^^&",

		metaInfo: PremiumUser{phone: 1234567890,
			address: "123 Main St",
			plan:    "Premium",
			website: "example.com"}}
	fmt.Printf("%+v", jon)
}
