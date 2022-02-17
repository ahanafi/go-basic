package main

// Import from package helper
import (
	"fmt"
	"go-basic/helper"
)

func main() {
	helper.SayHello("Ahmad")
	// helper.sayGoodBye("Hanafi") <-- It will error

	fmt.Println("App name:", helper.AppName)
	// fmt.Println("App version:", helper.version) <-- It will error
}
