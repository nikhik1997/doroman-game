package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Game represents the main game structure
type Game struct {
	Doroman1 *Doroman // First Doroman character
	Doroman2 *Doroman // Second Doroman character
}

// NewGame creates a new instance of Game with initialized Doroman characters
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	return &Game{
		Doroman1: NewDoroman("Doroman1", 100, 15), // Customize name, health, and attack power as needed
		Doroman2: NewDoroman("Doroman2", 100, 12), // Customize name, health, and attack power as needed
	}
}

// Start begins the game loop
func (g *Game) Start() {
	for {
		if !g.Doroman1.CanHit() && !g.Doroman2.CanHit() {
			fmt.Println("Both Doroman characters are knocked out. It's a draw!")
			break
		}

		if g.Doroman1.CanHit() {
			g.Doroman1.Hit(g.Doroman2)
			if g.Doroman2.Health <= 0 {
				break // Exit the loop if Doroman2 is knocked out
			}
		}

		if g.Doroman2.CanHit() {
			g.Doroman2.Hit(g.Doroman1)
			if g.Doroman1.Health <= 0 {
				break // Exit the loop if Doroman1 is knocked out
			}
		}
	}

	// Determine the winner
	var winner *Doroman
	if g.Doroman1.Health <= 0 && g.Doroman2.Health <= 0 {
		fmt.Println("Both Doroman characters are knocked out. It's a draw!")
	} else if g.Doroman1.Health <= 0 {
		winner = g.Doroman2
	} else {
		winner = g.Doroman1
	}

	// Celebrate the winner
	if winner != nil {
		winner.Celebrate()
	}
}
