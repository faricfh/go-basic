package helper

import "fmt"

var version = "1.0.0"
var AppName = "Go Basic"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Sample() {
	sayGoodbye("John")
	fmt.Println(version)
}
