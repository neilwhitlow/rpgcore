package stats_test

import (
	"runtime"
	"strconv"
	"testing"

	lhm "github.com/neilwhitlow/collections/linkedhashmap"
	"github.com/neilwhitlow/rpgcore/dice"
	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGetPrimeAbilitiesMap(t *testing.T) {
	t.Run("dummy", func(t *testing.T) {
		actual := stats.GetPrimeAbilityDefinitions()
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
			primeAbilities := stats.GetPrimeAbilityDefinitions()

			strength := primeAbilities.Get("STR")
			strength.InitialScore = test.init
			strength.ModifierBonus = test.bonus
			actual := strength.GetModifier()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestGetPrimeAbilityScoresAndMods(t *testing.T) {
	diceRoller := dice.NewRoller()
	t.Run("dummy", func(t *testing.T) {
		abilities := stats.GetPrimeAbilityDefinitions()
		if abilities == nil {
			t.Errorf("Expected abilities map, got nil instead)")
		}
		// for _, ability := range abilities {
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
		abilities := stats.GetCalculatedAbilityDefinitions()
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
		p := stats.Ability{}
		p.Abbreviation = "FAKE"
		p.InitialScore = 13

		primes := lhm.New[string, stats.Ability]()
		primes.Put(p.Abbreviation, p)

		c := stats.Ability{}
		c.PrimeModKey = "FAKE"
		result := c.GetScoreFromPrimeMod(primes)
		if result != 1 {
			t.Errorf("Expected 1, got %d instead)", result)
		}
	})
}

func TestGetScoreFromPrimeMultiplier(t *testing.T) {
	p := stats.Ability{}
	p.InitialScore = 10

	primes := lhm.New[string, stats.Ability]()
	primes.Put(p.Abbreviation, p)

	c := stats.Ability{}
	c.PrimeScoreMultiplier = 6

	for i := 1; i <= 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.GetScoreFromPrimeMultiple(primes)
			if result > 60 || result < 6 {
				t.Errorf("Expected range of 6-60, got %d instead)", result)
			}
		})
	}
}

func TestGetScoreFromMultiplier(t *testing.T) {
	p := stats.Ability{}
	p.Abbreviation = "STR"
	p.InitialScore = 15
	primes := lhm.New[string, stats.Ability]()
	primes.Put(p.Abbreviation, p)

	c := stats.Ability{}
	c.PrimeModKey = "STR"

	t.Run("STRMOD", func(t *testing.T) {
		c.InitialScore = c.GetScoreFromPrimeMod(primes)
		result := c.GetScore()
		if result != 2 {
			t.Errorf("Expected 2, got %d instead)", result)
		}
	})
}
