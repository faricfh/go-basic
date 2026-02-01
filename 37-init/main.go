package main

import (
	"fmt"
	"go-basic/37-init/database"
	_ "go-basic/37-init/internal" // blank identifier
)

func main() {
	fmt.Println(database.GetConnection())
}
