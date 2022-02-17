package main

/**
* Reference: https://pkg.go.dev/container/list
*/

import (
	"container/list"
	"fmt"
)

func main() {
	// Double link list
	fruits := list.New()
	fruits.PushBack("Apple")
	fruits.PushBack("Orange")
	fruits.PushBack("Banana")
	fruits.PushFront("Grape")

	fmt.Println("First fruit is: ", fruits.Front().Value)
	fmt.Println("Last fruit is: ", fruits.Back().Value)

	// Iterate from Front - End
	fmt.Println("Iterate from Front -> End")
	for fruit := fruits.Front(); fruit != nil; fruit = fruit.Next() {
		fmt.Println(fruit.Value)
	}

	// Iterate from the last - first
	fmt.Println("Iterate from Last -> Front")
	for fruit := fruits.Back(); fruit != nil; fruit = fruit.Prev() {
		fmt.Println(fruit.Value)
	}
}
