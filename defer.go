package main

import "fmt"

func logging()  {
	fmt.Println("Done calling function (FROM LOGGING)")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Running app")
	result := 10 / value
	fmt.Println("Result =", result)
}

func main() {
	runApp(10)

	// runApp(0) <-- Will error
}
