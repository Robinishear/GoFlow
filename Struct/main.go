package main

import "fmt"

type user struct {
	name     string
	email    string
	password string
}

func main() {
	jon := user{name: "jon", email: "jon@gmai.com", password: "12tgjyyyyyu8888;;l7itjtj3"} // name: is call key and value is

	fmt.Printf("%+v", jon)
}
