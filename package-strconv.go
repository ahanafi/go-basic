package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true1212")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error:", err.Error())
	}

	value := strconv.FormatInt(10000000, 10) // format as decimal base number (10)
	fmt.Println(value)
}
