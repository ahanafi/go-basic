package main

import (
	"fmt"
	"time"
)

/**
 * Reference: https://pkg.go.dev/time
 */


func main() {
	now := time.Now() // Get current time

	fmt.Println(now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())

	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Seconds:", now.Second())
	fmt.Println("Nano Seconds:", now.Nanosecond())

	utcTime := time.Date(2022, 02, 18, 7, 22, 32, 33, time.UTC)
	localTime := time.Date(2022, 02, 18, 7, 22, 32, 33, time.Local)
	fmt.Println("UTC Time:", utcTime)
	fmt.Println("Local Time:", localTime)

	// Reference format time: https://go.dev/src/time/format.go
	formatTime := "2006-01-02"
	parsedTime, _ := time.Parse(formatTime, "2022-02-18")
	parsedTime2, _ := time.Parse(formatTime, "1992-03-23")

	fmt.Println("Parsed time:", parsedTime)
	fmt.Println("Parsed time2:", parsedTime2)
}
