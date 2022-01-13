package main

import "fmt"

func main() {
	var name string // should be type data of variable if not assigning value

	name = "Ahmad Hanafi"
	fmt.Println(name)

	name = "Aris Munandar" // re-assign new value to variable name
	fmt.Println(name)

	// Not should type data type of variable when assign value
	var friendName = "Sugeng Sungkono"
	fmt.Println(friendName)

	var friendAge = 23
	fmt.Println(friendAge)

	country := "Indonesia" // Without 'var' for initialize new variable
	fmt.Println(country)

	country = "Turkey"
	fmt.Println(country)

	// Multiple declare variables
	var (
		firstName = "Ahmad"
		lastName  = "Hanafi"
	)

	fmt.Println("First name is = ", firstName)
	fmt.Println("last name is = ", lastName)
}
