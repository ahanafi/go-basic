package main

import (
	"fmt"
	"reflect"
)

/**
 * Reference: https://pkg.go.dev/reflect
 */

type Product struct {
	Name string `required:"true" max:"10"`
	Price int `required:"true"`
}

// Example validation method using reflect api
func isValid(data interface{}) bool  {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 7500000,
	}

	productType := reflect.TypeOf(product)
	fmt.Println("At product struct there are", productType.NumField(), "field")
	fmt.Println("First field is:", productType.Field(0).Name)
	fmt.Println("Second field is:", productType.Field(1).Name)

	// Get tag
	fmt.Println(productType.Field(0).Tag.Get("required"))
	fmt.Println(productType.Field(0).Tag.Get("max"))
	fmt.Println(productType.Field(0).Tag.Get("min")) // <-- tag doesn't exist will return empty string
}
