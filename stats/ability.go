package stats

import (
	_ "embed"
	"encoding/json"

	lhm "github.com/neilwhitlow/collections/linkedhashmap"
	"github.com/neilwhitlow/rpgcore/dice"
)

type Ability interface {
	GetName() string
	GetAbbreviation() string
	GetScore() int
}

type PrimeAbility struct {
	AbilityCore
}

type CalculatedAbility struct {
	AbilityCore
}

type AbilityCore struct {
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

func (a PrimeAbility) GetModifier() int {
	return GetScoreModifier(GetModifierMap(), a.GetScore()) + a.ModifierBonus
}

func (a PrimeAbility) GetScore() int {
	return a.InitialScore + a.ScoreBonus
}

func (a PrimeAbility) GetAbbreviation() string {
	return a.Abbreviation
}

func (a PrimeAbility) GetName() string {
	return a.Name
}

func (a CalculatedAbility) GetScore() int {
	return a.InitialScore + a.ScoreBonus
}

func (a CalculatedAbility) GetAbbreviation() string {
	return a.Abbreviation
}

func (a CalculatedAbility) GetName() string {
	return a.Name
}

//go:embed data/prime_abilities.json
var primeAbilityDefinitions []byte

func GetPrimeAbilitiesMap() *lhm.LinkedHashMap[string, PrimeAbility] {
	//a := make(map[string]PrimeAbility)

	abilities := lhm.New[string, PrimeAbility]()

	a := make([]PrimeAbility, 8)
	_ = json.Unmarshal(primeAbilityDefinitions, &a)
	for _, item := range a {
		abilities.Put(item.Abbreviation, item)
	}

	return abilities
}

//go:embed data/calculated_abilities.json
var calculatedAbilityDefinitions []byte

func GetCalculatedAbilitiesMap() *lhm.LinkedHashMap[string, CalculatedAbility] {
	abilities := lhm.New[string, CalculatedAbility]()

	a := make([]CalculatedAbility, 4)
	_ = json.Unmarshal(calculatedAbilityDefinitions, &a)
	for _, item := range a {
		abilities.Put(item.Abbreviation, item)
	}
	return abilities
}

func (ca CalculatedAbility) GetScoreFromPrimeMod(pa PrimeAbility) int {
	return pa.GetModifier()
}

func (ca CalculatedAbility) GetScoreFromPrimeMultiple(pa PrimeAbility) int {
	diceRoller := dice.NewDiceRoller()
	return diceRoller.RollMany(pa.GetScore(), dice.DieType(ca.PrimeScoreMultiplier))
}
