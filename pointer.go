package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Bandung",
		Province: "West Java",
		Country:  "Indonesia",
	}

	// address2 := address1  <--  Assign (Copy) the value (Pass the value)
	address2 := &address1 // <--  Pass value by reference to the address1 (Pointer)
	address2.City = "Cirebon"

	// *address2 = &Address{ <-- If we use this, all variable that refer to address1 will reference to address2
	// Will create new address
	address2 = &Address{
		City:     "Surabaya",
		Province: "East Java",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
}
