package main

import (
	"datafactory/generators"
	"fmt"
)

func main() {

	ints := generators.GenerateIntegers(10, 9, 90, "")
	fmt.Println(ints)

}
