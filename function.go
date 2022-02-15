package main

import "fmt"

// Declare function sayHello
func sayHello()  {
	fmt.Println("Hello World!")
}

// Function with parameter
func sayHelloTo(name string) {
	fmt.Println("Hello", name,"!")
}

// Function with return value
func sum(a int, b int) int {
	sum := a + b
	return sum
}

// Function with return multiple value
func getFullName() (string, string) {
	return "Ahmad", "Hanafi"
}

func friendsName() (string, string, string) {
	return "Deni", "Sugeng", "Aris"
}

// Function named return value
func getFullNameWithNamedReturnValue() (firstName string, lastName string) {
	firstName = "Ahmad"
	lastName = "Hanafi"

	return
}

func main()  {
	sayHello() // <-- Calling function

	sayHelloTo("Ahmad") // <-- Calling function and send 'Ahmad' as parameter
	sayHelloTo("Hanafi")

	result := sum(10, 50)
	fmt.Println("10 + 50 = ", result)

	firstName, lastName := getFullName()
	fmt.Println("Firstname is", firstName)
	fmt.Println("Last is", lastName)

	deny, sugeng, _ := friendsName() // <-- Ignoring last value with underscore ( _ )
	fmt.Println("First :", deny)
	fmt.Println("Second :", sugeng)

	ahmad, hanafi := getFullNameWithNamedReturnValue()
	fmt.Println("First name :", ahmad)
	fmt.Println("Last name :", hanafi)
}