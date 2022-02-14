package main

import "fmt"

func main() {
	name := "Hanafi"

	switch name {
	case "Hanafi":
		fmt.Println("Hallo Hanafi!")
	case "Ahmad":
		fmt.Println("Hai Ahmad!")
	default:
		fmt.Println("Boleh kenalan?")
	}

	// Switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name more than 5 character")
	case false:
		fmt.Println("Name equals or under 5 character")
	}

	// Switch without condition
	lengtName := len(name)
	switch {
	case lengtName > 10:
		fmt.Println("Name more than 10 character")
	case lengtName < 5:
		fmt.Println("Name under 5 character")
	default:
		fmt.Println("Name was correctly")

	}
}
