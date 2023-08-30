package main

import (
	"fmt"
	"math/rand"
)

// Game represents the state of the numbers printing game.
type Game struct {
	// The current number to be printed.
	CurrentNumber int
	// The number of times the current number has been printed.
	CurrentNumberCount int
	// The maximum number that can be printed.
	MaxNumber int
}

// NewGame creates a new game with a random starting number.
func NewGame(maxNumber int) Game {
	return Game{
		CurrentNumber: rand.Intn(maxNumber),
		CurrentNumberCount: 0,
		MaxNumber: maxNumber,
	}
}

// Play a turn of the game.
func (g *Game) Play() bool {
	// Check if the current number has been printed enough times.
	if g.CurrentNumberCount >= g.MaxNumber {
		return false
	}

	// Print the current number.
	fmt.Println(g.CurrentNumber)

	// Increment the number of times the current number has been printed.
	g.CurrentNumberCount++

	// Generate a new random number.
	g.CurrentNumber = rand.Intn(g.MaxNumber)

	return true
}

// main is the entry point for the program.
func main() {
	// Create a new game.
	g := NewGame(10)

	// Play the game until the game is over.
	for g.Play() {
	}
}