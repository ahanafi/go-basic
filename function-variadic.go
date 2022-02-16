package main

import "fmt"

// Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(10, 20, 30, 40)
	fmt.Println(result)

	// Put slice variable to variadic function
	sliceNumbers := []int{10, 10, 20, 30, 90}
	sumSlice := sumAll(sliceNumbers...)
	fmt.Println("Total slice:", sumSlice)
}
