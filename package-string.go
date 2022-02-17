package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "       Ahmad Hanafi      "

	// Trim => Remove space from start and end of string
	textTrimmed := strings.Trim(text, " ")

	// strings.Contains(s, substr) <-- Check if contains word in a string
	fmt.Println(strings.Contains(text, "Ahmad"))
	fmt.Println(strings.Contains(text, "Hanif"))

	// Split => to split string by separator
	fmt.Println(strings.Split(textTrimmed, " "))

	// ToLower => Convert to lowercase
	fmt.Println(strings.ToLower(textTrimmed))

	// ToUpper => Convert to uppercase
	fmt.Println(strings.ToUpper(textTrimmed))

	// ToTitle => Convert to unicode title case
	fmt.Println(strings.ToTitle(textTrimmed))

	// Replace || Replace All -> to replace string
	fmt.Println(strings.Replace(textTrimmed, "Hanafi", "Hanif", 1))
}