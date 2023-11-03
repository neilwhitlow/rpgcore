package stats_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/neilwhitlow/rpgcore/dice"
	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGetPrimeAbilitiesMap(t *testing.T) {
	t.Run("dummy", func(t *testing.T) {
		actual := stats.GetPrimeAbilitiesMap()
		assert.NotNil(t, actual, "Expected abilities map, got nil instead)")
	})
}

func TestGetPrimeAbilityModifier(t *testing.T) {

	tests := map[string]struct {
		init     int
		bonus    int
		expected int
	}{
		"3":          {init: 18, bonus: 0, expected: 3},
		"3WithBonus": {init: 16, bonus: 2, expected: 4},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			primeAbilities := stats.GetPrimeAbilitiesMap()

			strength := primeAbilities.Get("STR")
			strength.InitialScore = test.init
			strength.ModifierBonus = test.bonus
			actual := strength.GetModifier()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestGetPrimeAbilityScoresAndMods(t *testing.T) {
	diceRoller := dice.NewDiceRoller()
	t.Run("dummy", func(t *testing.T) {
		abilities := stats.GetPrimeAbilitiesMap()
		if abilities == nil {
			t.Errorf("Expected abilities map, got nil instead)")
		}
		//for _, ability := range abilities {
		for kvp := abilities.First(); kvp != nil; kvp = kvp.Next() {
			rollResult, err := diceRoller.RollAndDropLowest(4, dice.D6)
			kvp.Value.InitialScore = rollResult
			if err != nil {
				t.Errorf("Error rolling dice)")
			}
			t.Logf("%s: %d", kvp.Value.Name, kvp.Value.GetScore())
		}
	})
	runtime.GC()
}

func TestGetCalculatedAbilitiesMap(t *testing.T) {
	t.Run("dummy", func(t *testing.T) {
		abilities := stats.GetCalculatedAbilitiesMap()
		if abilities == nil {
			t.Errorf("Expected abilities map, got nil instead)")
		}
		for kvp := abilities.First(); kvp != nil; kvp = kvp.Next() {
			ability := kvp.Value
			t.Logf("%s: %s", ability.Name, ability.Abbreviation)
		}
	})
}

func TestGetScoreModifierOutOfBounds(t *testing.T) {
	t.Run("dummy", func(t *testing.T) {
		result := stats.GetScoreModifier(stats.GetModifierMap(), 100)
		if result != -999 {
			t.Errorf("Expected not found large negative indicator, got a potential value instead)")
		}
	})
}

func TestGetScoreFromPrimeMod(t *testing.T) {
	t.Run("dummy", func(t *testing.T) {

		p := stats.PrimeAbility{}
		p.InitialScore = 13

		c := stats.CalculatedAbility{}
		result := c.GetScoreFromPrimeMod(p)
		if result != 1 {
			t.Errorf("Expected 1, got %d instead)", result)
		}
	})
}

func TestGetScoreFromPrimeMultiplier(t *testing.T) {
	p := stats.PrimeAbility{}
	p.InitialScore = 10

	c := stats.CalculatedAbility{}
	c.PrimeScoreMultiplier = 6

	for i := 1; i <= 100; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := c.GetScoreFromPrimeMultiple(p)
			if result > 60 || result < 6 {
				t.Errorf("Expected range of 6-60, got %d instead)", result)
			}
		})
	}
}

func TestGetScoreFromMultiplier(t *testing.T) {
	p := stats.PrimeAbility{}
	p.Abbreviation = "STR"
	p.InitialScore = 15

	c := stats.CalculatedAbility{}
	c.PrimeModKey = "STR"

	t.Run("STRMOD", func(t *testing.T) {
		c.InitialScore = c.GetScoreFromPrimeMod(p)
		result := c.GetScore()
		if result != 2 {
			t.Errorf("Expected 2, got %d instead)", result)
		}
	})
}
