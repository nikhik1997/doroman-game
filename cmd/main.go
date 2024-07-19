package main

import (
	"math/rand"
	"time"

	"github.com/nikhik1997/doroman-game/internal/game/internal/game"
	//"github.com/nikhik1997/doroman-game/internal/game"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Initialize and start the game
	g := game.NewGame()
	g.Start()
}
