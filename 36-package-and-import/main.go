package main

import (
	"fmt"
	"go-basic/36-package-and-import/helper"
)

func main() {
	result := helper.SayHello("John")
	fmt.Println(result)

	fmt.Println(helper.AppName)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodbye("John"))
}
