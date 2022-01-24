package main

import "fmt"

func main() {
	type email string

	var myEmail email = "ahanafi.id@gmail.com"

	fmt.Println("Let's get in touch with me in ", myEmail)
}
