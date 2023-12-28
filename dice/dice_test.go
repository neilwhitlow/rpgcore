package dice_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/dice"
)

type rollDiceTypeTestDefinition struct {
	diceType      dice.DieType
	numberOfRolls int
	minExpected   int
	maxExpected   int
}

func TestRoll(t *testing.T) {
	tests := map[string]rollDiceTypeTestDefinition{
		"3-sided":  {diceType: dice.D3, minExpected: 1, maxExpected: 3},
		"4-sided":  {diceType: dice.D4, minExpected: 1, maxExpected: 4},
		"6-sided":  {diceType: dice.D6, minExpected: 1, maxExpected: 6},
		"8-sided":  {diceType: dice.D8, minExpected: 1, maxExpected: 8},
		"10-sided": {diceType: dice.D10, minExpected: 1, maxExpected: 10},
		"12-sided": {diceType: dice.D12, minExpected: 1, maxExpected: 12},
		"20-sided": {diceType: dice.D20, minExpected: 1, maxExpected: 20},
	}

	for i := 1; i <= 1000; i++ {
		for name, test := range tests {
			dr := dice.NewRoller()
			t.Run(name, func(t *testing.T) {
				result := dr.RollOnce(test.diceType)
				if result < test.minExpected || result > test.maxExpected {
					t.Errorf("Roll(%v) = %d; want %d-%d", test.diceType, result, test.minExpected, test.maxExpected)
				}
			})
		}
	}
}

func TestRollMany(t *testing.T) {
	tests := map[string]rollDiceTypeTestDefinition{
		"6-sided":             {diceType: dice.D6, minExpected: 3, maxExpected: 18, numberOfRolls: 3},
		"10-sided":            {diceType: dice.D10, minExpected: 10, maxExpected: 100, numberOfRolls: 10},
		"20-sided-zero-times": {diceType: dice.D20, minExpected: 0, maxExpected: 0, numberOfRolls: 0},
	}

	for i := 1; i <= 10; i++ {
		for name, test := range tests {
			dr := dice.NewRoller()
			t.Run(name, func(t *testing.T) {
				result := dr.RollMany(test.numberOfRolls, test.diceType)
				if result < test.minExpected || result > test.maxExpected {
					t.Errorf("Roll(%v) = %d; want %d-%d", test.diceType, result, test.minExpected, test.maxExpected)
				}
			})
		}
	}
}

func TestRollAndDropLowest(t *testing.T) {
	tests := map[string]rollDiceTypeTestDefinition{
		"6-sided":             {diceType: dice.D6, minExpected: 3, maxExpected: 18, numberOfRolls: 4},
		"10-sided":            {diceType: dice.D8, minExpected: 7, maxExpected: 56, numberOfRolls: 7},
		"20-sided-zero-times": {diceType: dice.D20, minExpected: 0, maxExpected: 0, numberOfRolls: 0},
	}

	for i := 1; i <= 5; i++ {
		for name, test := range tests {
			dr := dice.NewRoller()
			t.Run(name, func(t *testing.T) {
				result, err := dr.RollAndDropLowest(test.numberOfRolls, test.diceType)
				if err != nil {
					if test.numberOfRolls > 0 {
						t.Errorf("%v", err)
					}
				}
				if result < test.minExpected || result > test.maxExpected {
					t.Errorf("Roll(%v) = %d; want %d-%d", test.diceType, result, test.minExpected, test.maxExpected)
				}
			})
		}
	}
}
