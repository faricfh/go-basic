package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println("Value32:", value32)
	fmt.Println("Value64:", value64)
	fmt.Println("Value16:", value16)

	var name = "John Doe"
	var j = name[0]
	var jString = string(j)

	fmt.Println("Name:", name)
	fmt.Println("First character (byte):", j)
	fmt.Println("First character (string):", jString)
}
