package game

import (
	"fmt"
	"math/rand"
)

// Doroman represents a character in the game
type Doroman struct {
	Name        string // Name of the Doroman character
	Health      int    // Health points of the Doroman
	AttackPower int    // Power of each attack by the Doroman
	Points      int    // Accumulated points of the Doroman
}

// NewDoroman creates a new instance of Doroman with default values
func NewDoroman(name string, health, attackPower int) *Doroman {
	return &Doroman{
		Name:        name,
		Health:      health,
		AttackPower: attackPower,
		Points:      0,
	}
}

// Hit simulates one Doroman hitting another
func (d *Doroman) Hit(other *Doroman) {
	hitValue := rand.Intn(d.AttackPower) + 1 // Random hit value based on Doroman's attack power
	fmt.Printf("%s hits %s with %d points!\n", d.Name, other.Name, hitValue)
	other.Health -= hitValue
	if other.Health <= 0 {
		other.Health = 0
		fmt.Printf("%s has been knocked out!\n", other.Name)
		d.Points += 10 // Doroman earns points when opponent is knocked out
	}
}

// CanHit checks if the Doroman can still hit
func (d *Doroman) CanHit() bool {
	return d.Health > 0
}

// Celebrate performs a celebration dance when the Doroman wins
func (d *Doroman) Celebrate() {
	fmt.Printf("%s wins with %d points and starts dancing!\n", d.Name, d.Points)
}
