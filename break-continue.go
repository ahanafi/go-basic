package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			// Break -> to stop or exit loops
			break
		}

		fmt.Println("Index at ", i)
	}

	println("======================")

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue // -> Continue is for skipping to next loop
		}

		fmt.Println("Index at ", i)
	}
}
