package main

import "fmt"

func main()  {
	counter := 0

	increment := func() {
		fmt.Println("Counter +1")
		counter++ // <-- Accessed variable counter, to fix this, just create local variable or change name of variable
	}
	increment()
	fmt.Println(counter)
}