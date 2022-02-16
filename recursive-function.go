package main

import "fmt"

// Basic recursive with for loops
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func main() {
	_5factorial := factorialLoop(5)
	fmt.Println("5! =", _5factorial, " (USING LOOPS)")

	_5factorialWithRecursive := factorialRecursive(5)
	fmt.Println("5! =", _5factorialWithRecursive, " (USING RECURSIVE)")
}
