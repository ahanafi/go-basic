package main

import "fmt"

// Interface
type HasName interface {
	GetName() string
}

func sayHelloBro(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

// Struct
type Person struct {
	Name string
}

// Implement of the interface GetName
func (person Person) GetName() string {
	return person.Name
}

// Another Struct
type Animal struct {
	Name string
}

// Implement of the interface GetName
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ahmad Person
	ahmad.Name = "Ahmad"
	sayHelloBro(ahmad)

	wolf := Animal{
		Name: "Barry",
	}

	sayHelloBro(wolf)
}
