// Package character provides structs and operations for
// player and non player characters
package character

import (
	lhm "github.com/neilwhitlow/collections/linkedhashmap"
	"github.com/neilwhitlow/rpgcore/dice"
	"github.com/neilwhitlow/rpgcore/stats"
)

type Character struct {
	Name                string
	PrimeAbilities      *lhm.LinkedHashMap[string, stats.PrimeAbility]
	CalculatedAbilities *lhm.LinkedHashMap[string, stats.CalculatedAbility]
}

func New() (char Character) {
	c := Character{}
	c.Name = "New Character"
	c.PrimeAbilities = lhm.New[string, stats.PrimeAbility]()
	c.CalculatedAbilities = lhm.New[string, stats.CalculatedAbility]()
	return c
}

func GenerateCharacter() (char Character) {
	c := New()

	diceRoller := dice.NewDiceRoller()

	pa := stats.GetPrimeAbilitiesMap()
	if pa == nil {
		panic("Expected abilities map, got nil instead)")
	}
	for _, key := range pa.Keys() {
		a := pa.Get(key)
		rollResult, err := diceRoller.RollAndDropLowest(4, dice.D6)
		if err == nil {
			a.InitialScore = rollResult
			c.PrimeAbilities.Put(a.Abbreviation, a)
		}
	}

	ca := stats.GetCalculatedAbilitiesMap()
	if ca == nil {
		panic("Expected abilities map, got nil instead)")
	}

	for _, key := range ca.Keys() {
		a := ca.Get(key)
		initScore := 0
		if a.BaseInitialScore > 0 {
			initScore += a.BaseInitialScore
		}
		if a.PrimeScoreKey != "" {
			if a.PrimeScoreMultiplier != 0 {
				initScore += a.GetScoreFromPrimeMultiple(c.PrimeAbilities.Get((a.PrimeScoreKey)))
			} else {
				initScore += c.PrimeAbilities.Get(a.PrimeScoreKey).GetScore()
			}
		}
		if a.PrimeModKey != "" {
			initScore += c.PrimeAbilities.Get(a.PrimeModKey).GetModifier()
		}
		a.InitialScore = initScore
		c.CalculatedAbilities.Put(a.Abbreviation, a)
	}

	return c
}
