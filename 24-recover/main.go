package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Panic occurred :", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("Run App")
}

func main() {
	runApp(true)
	fmt.Println("Hello")
}
