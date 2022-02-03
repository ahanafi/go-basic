package main

import "fmt"

func main() {
	var months = [...]string {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	// take data at index 4 until 6 (4,5,6), but index at 7 not included
	// Value of firstSlice is May, June, July
	var firstSlice = months[4:7]
	fmt.Println("First Slice :" , firstSlice)
	fmt.Println("Length of firstSlice is ", len(firstSlice))
	fmt.Println("Capacity of firstSlice is ", cap(firstSlice))

	months[4] = "MEI" //Change array will be change in slice
	fmt.Println("First Slice: ", firstSlice)

	// Change value in slice also can change value of array
	firstSlice[1] = "JUNI"
	fmt.Println("MONTH : ", months)

	var secondSlice = months[10:] //take data from index 11 to last data
	fmt.Println("Second slice: " , secondSlice)

	// it will create new array because length of firstSlice is just 2
	// so, if we append 1 data, to existing slice, it will create new array
	var thirdSlice = append(secondSlice, "Hanafi")
	fmt.Println("Third slice: ", thirdSlice)

	// create new empty slice that not related other variable
	newSlice := make([]string, 2, 5) // 2 is length of slice, 5 is capacity
	newSlice[0] = "Ahmad"
	newSlice[1] = "Hanafi"
	fmt.Println("New slice: ", newSlice)
	fmt.Println("Length of New slice: ", len(newSlice))
	fmt.Println("Capacity of New slice: ", cap(newSlice))

	// Copying existing slice to new slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // Copying from newSlice to copySlice
	fmt.Println("CopySlice: ", copySlice)

	// Array vs Slice
	thisArray := [5]int{1, 2, 3, 4, 5} // Array is include length
	thisSlice := []int{1, 2, 3, 4, 5} // Slice is not include the length

	fmt.Println(thisArray)
	fmt.Println(thisSlice)

}
