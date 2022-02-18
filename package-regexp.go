package main

import (
	"fmt"
	"regexp"
)

/**
 * Reference: https://pkg.go.dev/regexp
 * Syntax reference: https://github.com/google/re2/wiki/Syntax
 */
func main() {
	regex := regexp.MustCompile("e[a-z]o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("emo"))
	fmt.Println(regex.MatchString("end"))
	fmt.Println(regex.MatchString("eDd"))

	result := regex.FindAllString("eko eka eki enda edo ego exi exa elo", 5)
	fmt.Println(result)
}
