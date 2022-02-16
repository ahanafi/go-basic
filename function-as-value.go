package main

import "fmt"

func getGoodBye(name string) string  {
	return "Good Bye " + name
}

func main() {
	// Function as value of variable
	sayGoodBye := getGoodBye
	result := sayGoodBye("Ahmad")
	fmt.Println(result)
}
