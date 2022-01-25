package main

import "fmt"

func main() {
	var (
		firstName = "Ahmad"
		lastName = "Hanafi"
	)

	var result = firstName == lastName

	fmt.Println("is", firstName," same with ",lastName," : ", result)

	var (
		val1 = 100
		val2 = 50
	)

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)


}
