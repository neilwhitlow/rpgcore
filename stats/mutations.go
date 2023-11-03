package stats

import (
	_ "embed"
	"encoding/json"
	"strings"

	"github.com/neilwhitlow/rpgcore/dice"
)

type RollChance struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type MutationDefinitions struct {
	MaxPhysicalRoll   int        `json:"maxPhysicalRoll"`
	MaxPsychicRoll    int        `json:"maxPsychicRoll"`
	PhysicalMutations []Mutation `json:"physicalMutations"`
	PsychicMutations  []Mutation `json:"psychicMutations"`
}
type Adjustments struct {
	Type          string `json:"type,omitempty"`
	AbilityKey    string `json:"abilityKey,omitempty"`
	InitialMod    int    `json:"initialMod,omitempty"`
	MinimumMod    int    `json:"minimumMod,omitempty"`
	ModifierBonus int    `json:"modBonus,omitempty"`
	ScoreBonus    int    `json:"scoreBonus,omitempty"`
}
type Mutation struct {
	Name                string        `json:"name"`
	RollChance          RollChance    `json:"rollChance"`
	HasMutationStrength bool          `json:"hasMutationStrength"`
	MutationStrength    int           `json:"mutationStrength,omitempty"`
	Adjustments         []Adjustments `json:"adjustments,omitempty"`
}

func (m Mutation) GetMutationStrengthModifier() int {
	return GetScoreModifier(GetModifierMap(), m.MutationStrength)
}

//go:embed data/mutations.json
var mutationDefinitions []byte

func GetMutationDefinitions() MutationDefinitions {
	m := MutationDefinitions{}
	_ = json.Unmarshal(mutationDefinitions, &m)
	return m
}

func (d *MutationDefinitions) GetRandomPhysicalMutation() Mutation {
	diceRoller := dice.NewDiceRoller()
	rollResult := diceRoller.RollFree(d.MaxPhysicalRoll)
	for _, physMut := range d.PhysicalMutations {
		if rollResult >= physMut.RollChance.From && rollResult <= physMut.RollChance.To {
			return physMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (d *MutationDefinitions) GetRandomPsychicMutation() Mutation {
	diceRoller := dice.NewDiceRoller()
	rollResult := diceRoller.RollFree(d.MaxPsychicRoll)
	for _, psychicMut := range d.PsychicMutations {
		if rollResult >= psychicMut.RollChance.From && rollResult <= psychicMut.RollChance.To {
			return psychicMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (d *MutationDefinitions) GetPhysicalMutation(mutationName string) Mutation {
	for _, physMut := range d.PhysicalMutations {
		if strings.ToUpper(physMut.Name) == strings.ToUpper(mutationName) {
			return physMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (d *MutationDefinitions) GetPsychicMutation(mutationName string) Mutation {
	for _, psychicMut := range d.PsychicMutations {
		if strings.ToUpper(psychicMut.Name) == strings.ToUpper(mutationName) {
			return psychicMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (m Mutation) RollMutationStrength() Mutation {
	if m.HasMutationStrength {
		diceRoller := dice.NewDiceRoller()
		rollResult, err := diceRoller.RollAndDropLowest(4, dice.D6)
		if err == nil {
			m.MutationStrength = rollResult
			return m
		}
	}
	return m
}
