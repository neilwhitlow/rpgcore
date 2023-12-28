// Package dice implements routines for generating randomizations
// for RPGs, most often expressed in terms and types of dice to be used
package dice

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/seehuhn/mt19937"
)

// DieType is a convenience function for popular dice notation
// d6 for a six sided die, d20 for a twenty sided die, etc.
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

// Roller is a random roller.
type Roller struct {
	rng *rand.Rand
}

// NewRoller returns a new random roller based on the
// 64bit Mersenne Twister pseudo random number generator
// from the package https://github.com/seehuhn/mt19937.
// The roller is seeded from the system time nanoseconds.
func NewRoller() Roller {
	dr := Roller{}
	dr.rng = rand.New(mt19937.New())
	dummy := time.Now().UnixNano()
	dr.rng.Seed(dummy)

	return dr
}

// RollOnce - a single roll from 1 to max, unconstrained by
// real-life polyhedral dice types.
func (dr *Roller) RollFree(max int) int {
	if max == 1 {
		return max
	}
	return dr.rng.Intn(max-1) + 1
}

// RollOnce- a single roll of the dice of the specified type.
// 1d6, 1d20, etc.
func (dr *Roller) RollOnce(dt DieType) int {
	return dr.rng.Intn(int(dt)-1) + 1
}

// RollMany - a convenience function for when a randomization
// requires multiple dice to be rolled.  For example: 3d6, or
// "percentile dice" (10d10).
func (dr *Roller) RollMany(numberOfDice int, dt DieType) int {
	var total int
	if numberOfDice <= 0 {
		return total
	}
	for i := 1; i <= numberOfDice; i++ {
		total += (dr.rng.Intn(int(dt)-1) + 1)
	}
	return total
}

// RollAndDropLowest - a convenience function similar to RollMany
// but automatically supports the practice found in many rulesets
// where the lowest roll is dropped to help protect against
// outrageously low results.  For example: 4d6 minus lowest die
// instead of a straight 3d6 roll for character attributes.
func (dr *Roller) RollAndDropLowest(numberOfDice int, dt DieType) (int, error) {
	var total int
	if numberOfDice <= 1 {
		return total, fmt.Errorf("NumberOfDice: %d, must be greater than one to drop lowest roll", numberOfDice)
	}
	rolls := make([]int, numberOfDice)
	for i := 0; i < numberOfDice; i++ {
		rolls[i] = (dr.rng.Intn(int(dt)) + 1)
	}

	sort.Ints(rolls)

	for _, v := range rolls[1:] {
		total += v
	}
	return total, nil
}

type RollChance struct {
	From int
	To   int
}
