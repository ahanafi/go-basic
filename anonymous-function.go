package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	// Anonymous function stored in variable
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Ahmad", blacklist)

	// Direct create anonymous function
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}