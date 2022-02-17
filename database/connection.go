package database

var connection string

// This function will be automatic call at first before calling other function
// It is same as constructor in other programming language
func init()  {
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}