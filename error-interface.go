package main

import (
	"errors"
	"fmt"
)

func divideNumber(number int, divide int) (int, error) {
	if divide == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return number / divide, nil
	}
}

func main() {
	// result, err := divideNumber(100, 0) <-- Will show error message
	result, err := divideNumber(100, 10)
	if err == nil {
		fmt.Println("100 / 10 =", result)
	} else {
		fmt.Println("ERROR:", err)
	}
}
