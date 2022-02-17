package main

import (
	"fmt"
	"go-basic/database"
	_ "go-basic/helper" // <-- Blank identifier
)

func main() {
	connection := database.GetConnection()
	fmt.Println(connection)
}
