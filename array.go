package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "Ahmad"
	names[1] = "Budi"
	names[2] = "Cika"
	names[3] = "Dedy"
	names[4] = "Eko"

	fmt.Println("Nama 1", names[0])
	fmt.Println("Nama 2", names[1])
	fmt.Println("Nama 3", names[2])
	fmt.Println("Nama 4", names[3])
	fmt.Println("Nama 5", names[4])
	// fmt.Println("Nama 6", names[5]) // errors, because out of index array

	var values = [4]int{
		80, 54, 66, 95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println("Length names is ", len(names), "items")
	fmt.Println("Length values is ", len(values), "items")
}
