package main

import "fmt"

// function type declaration
type Filter func(string) string

// Function as parameter
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Budi", spamFilter)

	sayHelloWithFilter("Anjing", spamFilter)
}
