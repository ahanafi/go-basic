package helper // <-- Package

import "fmt"

/**
* Access modifier => How the variable || function that can accessed out of package
* If variable/function name start from Uppercase, i.e SayHello, SayGoodBye <-- it can access from out of package
* But, if the variable/function name start from Lowercase, i.e. sayHello, sayGoodBy, etc. <-- it can`t access out of the package
*/

var AppName = "Go Basic" // <-- Can accessed
var version = 1.0		 // <-- Cannot accessed

func SayHello(name string) { 		// <-- Can accessed
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {		// <-- Cannot accessed
	fmt.Println("Good bye", name)
}