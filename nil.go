package main

import "fmt"

/**
 * Nil is same as Null in other programming language
 * Nil just can use for some data types such as Interface, function, map, slice, pointer, and channel.
 */

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person = NewMap("Ahmad")
	if person == nil {
		fmt.Println("Empty Name")
	} else {
		fmt.Println("Your name is", person["name"])
	}
}
