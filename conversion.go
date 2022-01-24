package main

import "fmt"

func main() {
	// Number conversion
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// String conversion
	var firstName = "Ahmad"
	var firstCharForFirstNameInByte = firstName[0]

	// Convert byte to string
	var firstChar = string(firstCharForFirstNameInByte)

	fmt.Println("First char from '",firstName,"' is ", firstChar)

}
