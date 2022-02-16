package main

import "fmt"

// Interface{}
func Oops(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Oops"
	}
}

func main() {
	oopsInt := Oops(1)
	oopsBool := Oops(2)
	oopsString := Oops(3)

	fmt.Println("OOPS 1 =", oopsInt)
	fmt.Println("OOPS 2 =", oopsBool)
	fmt.Println("OOPS 3 =", oopsString)

}