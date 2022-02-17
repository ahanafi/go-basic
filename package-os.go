package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // <-- Get argument
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname is", hostname)
	} else {
		fmt.Println("Hostname error:", err)
	}

	gopath := os.Getenv("GOPATH") // <-- Get data from Environment variable is OS
	fmt.Println("GOPATH :", gopath)
}
