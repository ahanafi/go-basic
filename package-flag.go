package main

import (
	"flag"
	"fmt"
)

func main() {
	/**
	* Command for running i.e =>  go run package-flag.go -hostname=alinx -username=root -password=kalkdsa
	*/


	host := flag.String("hostname", "localhost", "Put your hostname")
	user := flag.String("username", "root", "Put your username")
	pass := flag.String("password", "secret", "Put your password")

	flag.Parse() // <-- Parsing flag argument to get value from flag

	fmt.Println(*host, *user, *pass) // <-- All of variable from flag is defined as pointer

	fmt.Println("Hostname:", *host)
	fmt.Println("Username:", *host)
	fmt.Println("Password:", *host)
}
