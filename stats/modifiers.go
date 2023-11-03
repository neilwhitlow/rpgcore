// Package stats provides models and functions related to character stats
package stats

import (
	_ "embed"

	"github.com/jszwec/csvutil"
)

type AbilityModifer struct {
	Score    int `csv:"score"`
	Modifier int `csv:"mod"`
}

//go:embed data/score_modifiers.csv
var modifierData []byte

func GetModifierMap() map[int]int {
	var mods []AbilityModifer
	modiferMap := make(map[int]int)
	if err := csvutil.Unmarshal(modifierData, &mods); err != nil {
		return modiferMap
	}
	for _, m := range mods {
		modiferMap[m.Score] = m.Modifier
	}
	return modiferMap
}

func GetScoreModifier(modifiers map[int]int, score int) int {
	if modifier, ok := modifiers[score]; ok {
		return modifier
	}
	return -999
}
