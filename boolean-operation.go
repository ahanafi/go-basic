package main

import "fmt"

func main() {
	var (
		isFromJakarta = false
		isMarreid = false
		isAsian = true
		isIndonesian = true
		isMale = true
		isGraduatedFromCollege = true
	)

	fmt.Println(isFromJakarta && isIndonesian)
	fmt.Println(isMarreid && isMale)
	fmt.Println(isAsian && isIndonesian)
	fmt.Println(isMarreid && isGraduatedFromCollege)

}
