package main

import "fmt"

func random() interface{} {
	return "OK" // <-- String
	// return 100  <-- Int
	// return  true <-- Bool
}

func main() {
	var result = random()
	var resultString = result.(string) // <-- type assertion, convert from interface{} to string
	fmt.Println(resultString)

	// var resultInt = result.(int) // <-- It will be panic, because return type of random() is string, not int
	// fmt.Println(resultInt)

	// Safely type assertions with switch case
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Integer")
	case bool:
		fmt.Println("Value", value, "is Boolean")
	default:
		fmt.Println("Value", value, "is Unknown data types")
	}
}