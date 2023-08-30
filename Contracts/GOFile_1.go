package main

import (
	"fmt"
	"math/big"
)

// Add two numbers together
func add(a, b *big.Int) *big.Int {
	return a.Add(a, b)
}

func main() {
	// Create two big.Int numbers
	a := big.NewInt(10)
	b := big.NewInt(20)

	// Add the two numbers together
	c := add(a, b)

	// Print the result
	fmt.Println(c)
}