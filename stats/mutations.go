package stats

import (
	_ "embed"
	"encoding/json"
	"fmt"
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
	Type                   string `json:"type,omitempty"`
	AbilityKey             string `json:"abilityKey,omitempty"`
	InitialMod             int    `json:"initialMod,omitempty"`
	MinimumMod             int    `json:"minimumMod,omitempty"`
	ModifierBonus          int    `json:"modBonus,omitempty"`
	ScoreBonus             int    `json:"scoreBonus,omitempty"`
	AddMutationStrengthMod bool   `json:"addMutationStrengthMod,omitempty"`
}
type Mutation struct {
	Name                string        `json:"name"`
	RefinedName         string        `json:"refinedName,omitempty"`
	RollChance          RollChance    `json:"rollChance"`
	HasMutationStrength bool          `json:"hasMutationStrength"`
	MutationStrength    int           `json:"mutationStrength,omitempty"`
	Adjustments         []Adjustments `json:"adjustments,omitempty"`
	Refinements         []string      `json:"refinements,omitempty"`
	RefinedNameFormat   string        `json:"refinedNameFormat,omitempty"`
}

func (m Mutation) GetFinalName() string {
	return Coalesce(m.RefinedName, m.Name)
}

func (m Mutation) RollRefinedName() string {
	diceRoller := dice.NewRoller()
	rollResult := diceRoller.RollFree(len(m.Refinements))
	return fmt.Sprintf(m.RefinedNameFormat, m.Name, m.Refinements[rollResult-1])
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
	diceRoller := dice.NewRoller()
	rollResult := diceRoller.RollFree(d.MaxPhysicalRoll)
	for _, physMut := range d.PhysicalMutations {
		if rollResult >= physMut.RollChance.From && rollResult <= physMut.RollChance.To {
			return physMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (d *MutationDefinitions) GetRandomPsychicMutation() Mutation {
	diceRoller := dice.NewRoller()
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
		if strings.EqualFold(physMut.Name, mutationName) {
			return physMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

func (d *MutationDefinitions) GetPsychicMutation(mutationName string) Mutation {
	for _, psychicMut := range d.PsychicMutations {
		if strings.EqualFold(psychicMut.Name, mutationName) {
			return psychicMut.RollMutationStrength()
		}
	}
	return Mutation{}
}

// RollMutationStrength will check for the mutation strength property
// If the mutation has a strength, it is randomized and the mutation is
// returned with the value. If the mutation does not have strength, then
// the original mutation value is returned.
func (m Mutation) RollMutationStrength() Mutation {
	if m.HasMutationStrength {
		diceRoller := dice.NewRoller()
		rollResult, err := diceRoller.RollAndDropLowest(4, dice.D6)
		if err == nil {
			m.MutationStrength = rollResult
			return m
		}
	}
	return m
}
