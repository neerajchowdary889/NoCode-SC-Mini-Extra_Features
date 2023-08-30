package main

import (
	"fmt"
	"math/rand"
)

// Smart contract for a simple number game.
type Game struct {
	// The current number.
	Number int

	// The player's guess.
	Guess int

	// Whether the player has won.
	Won bool
}

// NewGame creates a new game with a random number.
func NewGame() *Game {
	return &Game{
		Number: rand.Intn(100),
	}
}

// Play makes a guess and returns whether the player has won.
func (g *Game) Play(guess int) bool {
	g.Guess = guess
	g.Won = g.Number == g.Guess
	return g.Won
}

// String returns a string representation of the game.
func (g *Game) String() string {
	return fmt.Sprintf("Game: number=%d, guess=%d, won=%t", g.Number, g.Guess, g.Won)
}

func main() {
	// Create a new game.
	g := NewGame()

	// Play the game.
	fmt.Println("Guess a number between 0 and 100:")
	guess := readInt()

	// Check if the player has won.
	if g.Play(guess) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}

	// Print the game state.
	fmt.Println(g)
}