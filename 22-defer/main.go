package main

import "fmt"

func logging() {
	fmt.Println("Logging")
}

func runApplication() {
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApplication()
}
