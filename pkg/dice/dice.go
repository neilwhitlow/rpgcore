// Package dice implements routines for generating randomizations
// for RPGs, most often expressed in terms and types of dice to be used
package dice

import (
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

// DieType is a convenience function for popular dice notation
// d6 for a six sided die, d20 for a twenty sided die, etc
type DieType int

const (
	D3  DieType = 3
	D4  DieType = 4
	D6  DieType = 6
	D8  DieType = 8
	D10 DieType = 10
	D12 DieType = 12
	D20 DieType = 20
)

// DiceRoller is a random roller
type DiceRoller struct {
	rng *rand.Rand
}

// NewDiceRoller returns a new random roller based on the
// 64bit Mersenne Twister pseudo random number generator
// from the package https://github.com/seehuhn/mt19937.
// The roller is seeded from the system time nanoseconds
func NewDiceRoller() *DiceRoller {
	dr := &DiceRoller{}
	dr.rng = rand.New(mt19937.New())
	dr.rng.Seed(time.Now().UnixNano())
	return dr
}

// RollOnce - a single roll of the dice of the specified type.
// 1d6, 1d20, etc
func (dr *DiceRoller) RollOnce(dt DieType) int {
	return dr.rng.Intn(int(dt)-1) + 1
}

// RollMany - a convenience function for when a randomization
// requires multiple dice to be rolled.  For example: 3d6, or
// "percentile dice" (10d10)
func (dr *DiceRoller) RollMany(numberOfDice int, dt DieType) int {
	var total int = 0
	if numberOfDice <= 0 {
		return total
	}
	for i := 1; i <= numberOfDice; i++ {
		total += (dr.rng.Intn(int(dt)-1) + 1)
	}
	return total
}

type RollChance struct {
	From int
	To   int
}
