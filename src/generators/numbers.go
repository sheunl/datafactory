package generators

import (
	"fmt"
	"math/rand"
)

// Function to generate integers 
func GenerateIntegers(size int, min int, max int, format string) []int {
	
	integers := []int{}

	for i:= 0; i < size; i++ {
		integers = append(integers, rand.Intn(max-min + 1) + min )
	}

	return integers
}