package main

import "fmt"

func main() {
	counter := 1

	// Basic For loops
	for counter <= 10 {
		fmt.Println("Loops at : ", counter)
		counter++
	}

	// For with statement
	for index := 1; index <= 30; index++ {
		fmt.Println("Index at : ", index)
	}

	// Print array or slice data with loops
	names := []string{"Ahmad", "Hanafi", "Reza", "Ade", "Yudi"}

	// Using for only (Without range)
	for i := 0; i < len(names); i++ {
		fmt.Println("Name at index ", i, " is ", names[i])
	}

	// Using for - range (With Index)
	for index, name := range names {
		fmt.Println("Name at ", index, "is ", name)
	}

	// Using for - range (Without Index)
	for _, name := range names {
		fmt.Println("Name is ", name)
	}

	person := make(map[string]string)
	person["name"] = "Ahmad"
	person["jobs"] = "Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
