package main

import "fmt"

func endApp() {
	fmt.Println("Done app")

	// Catch message from panic error
	message := recover() // <-- Recover should be store on defer function
	if message != nil {
		fmt.Println("ERROR MESSAGE: ",message)
	}
}

func runApplication(error bool) {
	defer endApp() // <-- endApp call as defer function
	if error {
		panic("APP ERROR!!!")
	}

	fmt.Println("Application running...")
}

func main() {
	runApplication(false)

	// runApplication(true) <-- Will Error
}
