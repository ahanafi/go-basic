package main

/**
 * Reference: https://pkg.go.dev/sort
 */

import (
	"fmt"
	"sort"
)

type People struct {
	Name string
	Age int
}

type PeopleSlice []People

func (value PeopleSlice) Len() int {
	return len(value)
}

func (value PeopleSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

// Switching || swap the value
func (value PeopleSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	peoples := []People {
		{"Hanafi", 30},
		{"Ahmad", 20},
		{"Dedi", 56},
		{"Eka", 24},
		{"Vika", 32},
		{"Zeke", 44},
		{"Budi", 17},
		{"Andre", 15},
	}

	fmt.Println("BEFORE SORT :", peoples)
	sort.Sort(PeopleSlice(peoples))
	fmt.Println("AFTER SORT :", peoples)
}
