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
	PhysicalMutations   []stats.Mutation
	PsychicMutations    []stats.Mutation
}

func New() (char Character) {
	c := Character{}
	c.Name = "New Character"
	c.PrimeAbilities = lhm.New[string, stats.PrimeAbility]()
	c.CalculatedAbilities = lhm.New[string, stats.CalculatedAbility]()
	c.PhysicalMutations = make([]stats.Mutation, 0)
	c.PsychicMutations = make([]stats.Mutation, 0)
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
		c.CalculatedAbilities.Put(a.Abbreviation, a)
	}

	mutationDefinitions := stats.GetMutationDefinitions()

	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetPhysicalMutation("Enhanced Accuracy"))

	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetRandomPhysicalMutation())
	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetRandomPhysicalMutation())
	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetRandomPhysicalMutation())
	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetRandomPhysicalMutation())
	c.PhysicalMutations = append(c.PhysicalMutations, mutationDefinitions.GetRandomPhysicalMutation())

	for _, physMut := range c.PhysicalMutations {
		for _, adjustment := range physMut.Adjustments {
			if adjustment.Type == "PrimeAbility" {
				ability := c.PrimeAbilities.Get(adjustment.AbilityKey)
				adjustmentCalc := Max(adjustment.MinimumMod, (adjustment.InitialMod - ability.GetModifier()))
				ability.InitialScore += adjustmentCalc
				c.PrimeAbilities.Put(ability.Abbreviation, ability)
			}
			if adjustment.Type == "CalculatedAbilityBonus" {
				ability := c.CalculatedAbilities.Get(adjustment.AbilityKey)
				if adjustment.ScoreBonus != 0 {
					ability.ScoreBonus += adjustment.ScoreBonus
					c.CalculatedAbilities.Put(ability.Abbreviation, ability)
				}
			}
		}
	}

	c.PsychicMutations = append(c.PsychicMutations, mutationDefinitions.GetRandomPsychicMutation())

	for _, psychicMut := range c.PsychicMutations {
		for _, adjustment := range psychicMut.Adjustments {
			if adjustment.Type == "PrimeAbility" {
				ability := c.PrimeAbilities.Get(adjustment.AbilityKey)
				adjustmentCalc := Max(adjustment.MinimumMod, (adjustment.InitialMod - ability.GetModifier()))
				ability.InitialScore += adjustmentCalc
				c.PrimeAbilities.Put(ability.Abbreviation, ability)
			}
			if adjustment.Type == "CalculatedAbilityBonus" {
				ability := c.CalculatedAbilities.Get(adjustment.AbilityKey)
				if adjustment.ScoreBonus != 0 {
					ability.ScoreBonus += adjustment.ScoreBonus
					c.CalculatedAbilities.Put(ability.Abbreviation, ability)
				}
			}
		}
	}

	for _, key := range ca.Keys() {
		a := c.CalculatedAbilities.Get(key)
		initScore := a.InitialScore
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

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
