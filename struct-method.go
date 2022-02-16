package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// Assign function || method to the struct
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "!, My name is", customer.Name)
}

func main() {
	var ahmad Customer
	ahmad.Name = "Ahmad"
	ahmad.Age  = 23
	ahmad.Address = "Bandung"
	ahmad.sayHi("Hanafi")
}
