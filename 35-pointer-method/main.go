package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// pointer in method no need use &
	man := Man{"John"}
	man.Married()

	fmt.Println(man)
}
