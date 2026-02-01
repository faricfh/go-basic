package main

import "fmt"

func main() {
	var score = 90
	var attendance = 85

	var passScore bool = score > 80
	var passAttendance bool = attendance > 80

	var passed bool = passScore && passAttendance

	fmt.Println(passed)
}
