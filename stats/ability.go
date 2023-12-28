package stats

import (
	_ "embed"
	"encoding/json"

	lhm "github.com/neilwhitlow/collections/linkedhashmap"
	"github.com/neilwhitlow/rpgcore/dice"
)

type AbilityDisplay interface {
	GetName() string
	GetAbbreviation() string
	GetScore() int
	GetModifier() int
}
type Ability struct {
	Name                  string `json:"name"`
	Abbreviation          string `json:"abbreviation"`
	PrimeModKey           string `json:"primeModKey,omitempty"`
	PrimeScoreKey         string `json:"primeScoreKey,omitempty"`
	PrimeScoreMultiplier  int    `json:"primeScoreMultiplier,omitempty"`
	BaseInitialScore      int    `json:"baseInitialScore,omitempty"`
	InitialScore          int
	ScoreBonus            int
	ModifierBonus         int
	HasBeenAssignedPoints bool
	MinimumInitialScore   int
	MaximumInitialScore   int
}

func (a Ability) GetScore() int {
	return a.InitialScore + a.ScoreBonus
}

func (a Ability) GetAbbreviation() string {
	return a.Abbreviation
}

func (a Ability) GetName() string {
	return a.Name
}

func (a Ability) GetModifier() int {
	return GetScoreModifier(GetModifierMap(), a.GetScore()) + a.ModifierBonus
}

//go:embed data/prime_abilities.json
var primeAbilityDefinitions []byte

func GetPrimeAbilityDefinitions() *lhm.LinkedHashMap[string, Ability] {
	abilities := lhm.New[string, Ability]()

	a := make([]Ability, 8)
	_ = json.Unmarshal(primeAbilityDefinitions, &a)
	for _, item := range a {
		abilities.Put(item.Abbreviation, item)
	}

	return abilities
}

//go:embed data/calculated_abilities.json
var calculatedAbilityDefinitions []byte

func GetCalculatedAbilityDefinitions() *lhm.LinkedHashMap[string, Ability] {
	abilities := lhm.New[string, Ability]()

	a := make([]Ability, 8)
	_ = json.Unmarshal(calculatedAbilityDefinitions, &a)
	for _, item := range a {
		abilities.Put(item.Abbreviation, item)
	}
	return abilities
}

// GetScoreFromPrimeMod
// Currently, this is a simple implementation to return the prime
// ability's modifier. Ability in method receiver should only ever be a Calculated Ability.
func (a Ability) GetScoreFromPrimeMod(primeAbilities *lhm.LinkedHashMap[string, Ability]) int {
	return primeAbilities.Get(a.PrimeModKey).GetModifier()
}

// GetScoreFromPrimeMultiple
// Roll a dice of a given type a number of times equal to the current final score
// of a Prime ability. Ability in method receiver should only ever be a Calculated Ability.
func (a Ability) GetScoreFromPrimeMultiple(primeAbilities *lhm.LinkedHashMap[string, Ability]) int {
	diceRoller := dice.NewRoller()
	primeScore := primeAbilities.Get(a.PrimeScoreKey).GetScore()
	return diceRoller.RollMany(primeScore, dice.DieType(a.PrimeScoreMultiplier))
}
