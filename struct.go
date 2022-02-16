package main

import "fmt"

// Define struct User
type User struct {
	Username, Email, Password string
	Age int
}

func main() {
	var user1 User
	// Set value struct user1
	user1.Username = "ahanafi"
	user1.Email = "ahanafi@gmail.com"
	user1.Password = "secret"
	user1.Age = 23

	// Get value from struct user1
	fmt.Println("Username: ", user1.Username)
	fmt.Println("Password: ", user1.Password)
	fmt.Println("Email   : ", user1.Email)
	fmt.Println("Age     : ", user1.Age," years old")

	user2 := User{
		Username: "ahmad",
		Password: "123456",
		Email: "ahmad@gmail.com",
		Age: 22,
	}
	fmt.Println(user2)
}
